package parser

// ASTNode is the interface for all AST nodes
type ASTNode interface {
	Type() string
}

// SelectStatement represents a SELECT query
type SelectStatement struct {
	Columns   []string
	TableName string
	Where     *WhereClause
}

func (s *SelectStatement) Type() string {
	return "SELECT"
}

// InsertStatement represents an INSERT query
type InsertStatement struct {
	TableName string
	Columns   []string
	Values    [][]interface{}
}

func (i *InsertStatement) Type() string {
	return "INSERT"
}

// UpdateStatement represents an UPDATE query
type UpdateStatement struct {
	TableName string
	SetClause map[string]interface{}
	Where     *WhereClause
}

func (u *UpdateStatement) Type() string {
	return "UPDATE"
}

// DeleteStatement represents a DELETE query
type DeleteStatement struct {
	TableName string
	Where     *WhereClause
}

func (d *DeleteStatement) Type() string {
	return "DELETE"
}

// CreateTableStatement represents a CREATE TABLE query
type CreateTableStatement struct {
	TableName string
	Columns   []ColumnDef
}

func (c *CreateTableStatement) Type() string {
	return "CREATE_TABLE"
}

// DropTableStatement represents a DROP TABLE query
type DropTableStatement struct {
	TableName string
}

func (d *DropTableStatement) Type() string {
	return "DROP_TABLE"
}

// ColumnDef represents a column definition
type ColumnDef struct {
	Name string
	Type string
}

// WhereClause represents a WHERE condition
type WhereClause struct {
	Conditions []Condition
}

// Condition represents a single condition in WHERE clause
type Condition struct {
	Column   string
	Operator string
	Value    interface{}
}

// Expression represents a generic expression
type Expression struct {
	Left     interface{}
	Operator string
	Right    interface{}
}
