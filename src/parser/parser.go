package parser

// Statement represents a parsed SQL statement
type Statement interface {
	Type() string
}

// SelectStatement represents a SELECT query
type SelectStatement struct {
	Table   string
	Columns []string
	Where   string
}

func (s *SelectStatement) Type() string {
	return "SELECT"
}

// Parser parses SQL statements
type Parser struct{}

// NewParser creates a new SQL parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a SQL statement (stub implementation)
func (p *Parser) Parse(sql string) (Statement, error) {
	// TODO: Implement actual SQL parsing
	return &SelectStatement{}, nil
}
