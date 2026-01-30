package executor

import (
	"fmt"

	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/index"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/parser"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/storage"
	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/transaction"
)

// ExecutionResult represents the result of a query execution
type ExecutionResult struct {
	Rows         []map[string]interface{}
	RowsAffected int
	Message      string
}

// Executor executes SQL queries
type Executor struct {
	tableManager *storage.TableManager
	indexManager *index.BTreeManager
	walManager   *transaction.WALManager
	lockManager  *transaction.LockManager
}

// NewExecutor creates a new query executor
func NewExecutor(
	tableManager *storage.TableManager,
	indexManager *index.BTreeManager,
	walManager *transaction.WALManager,
	lockManager *transaction.LockManager,
) *Executor {
	return &Executor{
		tableManager: tableManager,
		indexManager: indexManager,
		walManager:   walManager,
		lockManager:  lockManager,
	}
}

// Execute executes an AST node and returns the result
func (e *Executor) Execute(node parser.ASTNode) (*ExecutionResult, error) {
	switch stmt := node.(type) {
	case *parser.SelectStatement:
		return e.executeSelect(stmt)
	case *parser.InsertStatement:
		return e.executeInsert(stmt)
	case *parser.UpdateStatement:
		return e.executeUpdate(stmt)
	case *parser.DeleteStatement:
		return e.executeDelete(stmt)
	case *parser.CreateTableStatement:
		return e.executeCreateTable(stmt)
	case *parser.DropTableStatement:
		return e.executeDropTable(stmt)
	default:
		return nil, fmt.Errorf("unsupported statement type")
	}
}

// executeSelect executes a SELECT statement
func (e *Executor) executeSelect(stmt *parser.SelectStatement) (*ExecutionResult, error) {
	// Get table
	table, err := e.tableManager.GetTable(stmt.TableName)
	if err != nil {
		return nil, err
	}

	// Acquire read lock
	if err := e.lockManager.AcquireReadLock(stmt.TableName); err != nil {
		return nil, err
	}
	defer e.lockManager.ReleaseReadLock(stmt.TableName)

	// Scan table (stub - return empty result)
	tuples, err := table.Scan()
	if err != nil {
		return nil, err
	}

	// Convert tuples to result rows
	rows := make([]map[string]interface{}, 0)
	for range tuples {
		// TODO: Convert tuple to map based on schema and selected columns
		row := make(map[string]interface{})
		rows = append(rows, row)
	}

	return &ExecutionResult{
		Rows:    rows,
		Message: fmt.Sprintf("Selected %d rows from %s", len(rows), stmt.TableName),
	}, nil
}

// executeInsert executes an INSERT statement
func (e *Executor) executeInsert(stmt *parser.InsertStatement) (*ExecutionResult, error) {
	// Get table
	table, err := e.tableManager.GetTable(stmt.TableName)
	if err != nil {
		return nil, err
	}

	// Acquire write lock
	if err := e.lockManager.AcquireWriteLock(stmt.TableName); err != nil {
		return nil, err
	}
	defer e.lockManager.ReleaseWriteLock(stmt.TableName)

	// Log to WAL
	e.walManager.LogInsert(stmt.TableName, stmt.Values)

	// Insert tuples
	rowsAffected := 0
	for _, values := range stmt.Values {
		tuple := storage.Tuple{Values: values}
		if err := table.InsertTuple(tuple); err != nil {
			return nil, err
		}
		rowsAffected++
	}

	return &ExecutionResult{
		RowsAffected: rowsAffected,
		Message:      fmt.Sprintf("Inserted %d row(s) into %s", rowsAffected, stmt.TableName),
	}, nil
}

// executeUpdate executes an UPDATE statement
func (e *Executor) executeUpdate(stmt *parser.UpdateStatement) (*ExecutionResult, error) {
	// Get table
	_, err := e.tableManager.GetTable(stmt.TableName)
	if err != nil {
		return nil, err
	}

	// Acquire write lock
	if err := e.lockManager.AcquireWriteLock(stmt.TableName); err != nil {
		return nil, err
	}
	defer e.lockManager.ReleaseWriteLock(stmt.TableName)

	// Log to WAL
	e.walManager.LogUpdate(stmt.TableName, stmt.SetClause, stmt.Where)

	// TODO: Implement actual update logic
	rowsAffected := 0

	return &ExecutionResult{
		RowsAffected: rowsAffected,
		Message:      fmt.Sprintf("Updated %d row(s) in %s", rowsAffected, stmt.TableName),
	}, nil
}

// executeDelete executes a DELETE statement
func (e *Executor) executeDelete(stmt *parser.DeleteStatement) (*ExecutionResult, error) {
	// Get table
	_, err := e.tableManager.GetTable(stmt.TableName)
	if err != nil {
		return nil, err
	}

	// Acquire write lock
	if err := e.lockManager.AcquireWriteLock(stmt.TableName); err != nil {
		return nil, err
	}
	defer e.lockManager.ReleaseWriteLock(stmt.TableName)

	// Log to WAL
	e.walManager.LogDelete(stmt.TableName, stmt.Where)

	// TODO: Implement actual delete logic
	rowsAffected := 0

	return &ExecutionResult{
		RowsAffected: rowsAffected,
		Message:      fmt.Sprintf("Deleted %d row(s) from %s", rowsAffected, stmt.TableName),
	}, nil
}

// executeCreateTable executes a CREATE TABLE statement
func (e *Executor) executeCreateTable(stmt *parser.CreateTableStatement) (*ExecutionResult, error) {
	// Convert column definitions to schema
	columns := make([]storage.Column, len(stmt.Columns))
	for i, col := range stmt.Columns {
		columns[i] = storage.Column{
			Name: col.Name,
			Type: col.Type,
		}
	}

	schema := storage.Schema{Columns: columns}

	// Create table
	if err := e.tableManager.CreateTable(stmt.TableName, schema); err != nil {
		return nil, err
	}

	return &ExecutionResult{
		Message: fmt.Sprintf("Created table %s", stmt.TableName),
	}, nil
}

// executeDropTable executes a DROP TABLE statement
func (e *Executor) executeDropTable(stmt *parser.DropTableStatement) (*ExecutionResult, error) {
	// Drop table
	if err := e.tableManager.DropTable(stmt.TableName); err != nil {
		return nil, err
	}

	return &ExecutionResult{
		Message: fmt.Sprintf("Dropped table %s", stmt.TableName),
	}, nil
}
