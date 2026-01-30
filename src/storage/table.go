package storage

import (
	"fmt"
	"sync"
)

// Column represents a column in a table
type Column struct {
	Name string
	Type string // "int", "string", "float"
}

// Schema represents the schema of a table
type Schema struct {
	Columns []Column
}

// Tuple represents a row in a table
type Tuple struct {
	Values []interface{}
}

// Table represents a database table
type Table struct {
	Name       string
	Schema     Schema
	Pages      []PageID
	bufferPool *BufferPool
	mu         sync.RWMutex
}

// TableManager manages all tables in the database
type TableManager struct {
	tables     map[string]*Table
	bufferPool *BufferPool
	nextPageID PageID
	mu         sync.RWMutex
}

// NewTableManager creates a new table manager
func NewTableManager(bufferPool *BufferPool) *TableManager {
	return &TableManager{
		tables:     make(map[string]*Table),
		bufferPool: bufferPool,
		nextPageID: 0,
	}
}

// CreateTable creates a new table with the given schema
func (tm *TableManager) CreateTable(name string, schema Schema) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, exists := tm.tables[name]; exists {
		return fmt.Errorf("table %s already exists", name)
	}

	table := &Table{
		Name:       name,
		Schema:     schema,
		Pages:      make([]PageID, 0),
		bufferPool: tm.bufferPool,
	}

	// Allocate first page for the table
	firstPageID := tm.allocatePage()
	table.Pages = append(table.Pages, firstPageID)

	tm.tables[name] = table
	return nil
}

// DropTable removes a table from the database
func (tm *TableManager) DropTable(name string) error {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if _, exists := tm.tables[name]; !exists {
		return fmt.Errorf("table %s does not exist", name)
	}

	delete(tm.tables, name)
	return nil
}

// GetTable retrieves a table by name
func (tm *TableManager) GetTable(name string) (*Table, error) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	table, exists := tm.tables[name]
	if !exists {
		return nil, fmt.Errorf("table %s does not exist", name)
	}

	return table, nil
}

// ListTables returns all table names
func (tm *TableManager) ListTables() []string {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	names := make([]string, 0, len(tm.tables))
	for name := range tm.tables {
		names = append(names, name)
	}
	return names
}

// allocatePage allocates a new page
func (tm *TableManager) allocatePage() PageID {
	pageID := tm.nextPageID
	tm.nextPageID++
	return pageID
}

// InsertTuple inserts a tuple into a table
func (t *Table) InsertTuple(tuple Tuple) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	// Validate tuple matches schema
	if len(tuple.Values) != len(t.Schema.Columns) {
		return fmt.Errorf("tuple has %d values but schema has %d columns",
			len(tuple.Values), len(t.Schema.Columns))
	}

	// TODO: Implement actual tuple insertion into pages
	// For now, just a stub
	return nil
}

// Scan returns an iterator over all tuples in the table
func (t *Table) Scan() ([]Tuple, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	// TODO: Implement actual table scan
	// For now, return empty result
	return []Tuple{}, nil
}

// GetSchema returns the table's schema
func (t *Table) GetSchema() Schema {
	return t.Schema
}
