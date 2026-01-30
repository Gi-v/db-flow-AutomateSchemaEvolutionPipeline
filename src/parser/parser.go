package parser

import (
	"fmt"
	"strconv"
)

// Parser parses SQL statements into AST
type Parser struct {
	lexer   *Lexer
	curTok  Token
	peekTok Token
}

// NewParser creates a new parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse parses a SQL statement
func (p *Parser) Parse(input string) (ASTNode, error) {
	p.lexer = NewLexer(input)
	p.nextToken()
	p.nextToken()

	switch p.curTok.Type {
	case TokenSelect:
		return p.parseSelect()
	case TokenInsert:
		return p.parseInsert()
	case TokenUpdate:
		return p.parseUpdate()
	case TokenDelete:
		return p.parseDelete()
	case TokenCreate:
		return p.parseCreate()
	case TokenDrop:
		return p.parseDrop()
	default:
		return nil, fmt.Errorf("unexpected token: %s", p.curTok.Literal)
	}
}

// nextToken advances to the next token
func (p *Parser) nextToken() {
	p.curTok = p.peekTok
	p.peekTok = p.lexer.NextToken()
}

// expectToken checks if current token matches expected type
func (p *Parser) expectToken(t TokenType) error {
	if p.curTok.Type != t {
		return fmt.Errorf("expected %s, got %s", t, p.curTok.Type)
	}
	return nil
}

// parseSelect parses a SELECT statement
func (p *Parser) parseSelect() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	// Parse columns
	p.nextToken()
	if p.curTok.Type == TokenStar {
		stmt.Columns = []string{"*"}
		p.nextToken()
	} else {
		columns := []string{}
		for {
			if p.curTok.Type != TokenIdentifier {
				return nil, fmt.Errorf("expected column name")
			}
			columns = append(columns, p.curTok.Literal)
			p.nextToken()

			if p.curTok.Type != TokenComma {
				break
			}
			p.nextToken()
		}
		stmt.Columns = columns
	}

	// Parse FROM clause
	if err := p.expectToken(TokenFrom); err != nil {
		return nil, err
	}
	p.nextToken()

	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected table name")
	}
	stmt.TableName = p.curTok.Literal
	p.nextToken()

	// Parse optional WHERE clause
	if p.curTok.Type == TokenWhere {
		where, err := p.parseWhere()
		if err != nil {
			return nil, err
		}
		stmt.Where = where
	}

	return stmt, nil
}

// parseInsert parses an INSERT statement
func (p *Parser) parseInsert() (*InsertStatement, error) {
	stmt := &InsertStatement{}

	// Parse INTO
	p.nextToken()
	if err := p.expectToken(TokenInto); err != nil {
		return nil, err
	}
	p.nextToken()

	// Parse table name
	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected table name")
	}
	stmt.TableName = p.curTok.Literal
	p.nextToken()

	// Parse VALUES
	if err := p.expectToken(TokenValues); err != nil {
		return nil, err
	}
	p.nextToken()

	// Parse values (simplified - just one row)
	if err := p.expectToken(TokenLeftParen); err != nil {
		return nil, err
	}
	p.nextToken()

	values := []interface{}{}
	for p.curTok.Type != TokenRightParen {
		var val interface{}
		switch p.curTok.Type {
		case TokenNumber:
			num, _ := strconv.Atoi(p.curTok.Literal)
			val = num
		case TokenString:
			val = p.curTok.Literal
		default:
			return nil, fmt.Errorf("unexpected value type")
		}
		values = append(values, val)
		p.nextToken()

		if p.curTok.Type == TokenComma {
			p.nextToken()
		}
	}

	stmt.Values = [][]interface{}{values}
	return stmt, nil
}

// parseUpdate parses an UPDATE statement
func (p *Parser) parseUpdate() (*UpdateStatement, error) {
	stmt := &UpdateStatement{
		SetClause: make(map[string]interface{}),
	}

	// Parse table name
	p.nextToken()
	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected table name")
	}
	stmt.TableName = p.curTok.Literal
	p.nextToken()

	// Parse SET clause
	if err := p.expectToken(TokenSet); err != nil {
		return nil, err
	}
	p.nextToken()

	// Parse column = value pairs (simplified)
	for {
		if p.curTok.Type != TokenIdentifier {
			return nil, fmt.Errorf("expected column name")
		}
		col := p.curTok.Literal
		p.nextToken()

		if err := p.expectToken(TokenEqual); err != nil {
			return nil, err
		}
		p.nextToken()

		var val interface{}
		switch p.curTok.Type {
		case TokenNumber:
			num, _ := strconv.Atoi(p.curTok.Literal)
			val = num
		case TokenString:
			val = p.curTok.Literal
		default:
			return nil, fmt.Errorf("unexpected value type")
		}

		stmt.SetClause[col] = val
		p.nextToken()

		if p.curTok.Type != TokenComma {
			break
		}
		p.nextToken()
	}

	// Parse optional WHERE clause
	if p.curTok.Type == TokenWhere {
		where, err := p.parseWhere()
		if err != nil {
			return nil, err
		}
		stmt.Where = where
	}

	return stmt, nil
}

// parseDelete parses a DELETE statement
func (p *Parser) parseDelete() (*DeleteStatement, error) {
	stmt := &DeleteStatement{}

	// Parse FROM
	p.nextToken()
	if err := p.expectToken(TokenFrom); err != nil {
		return nil, err
	}
	p.nextToken()

	// Parse table name
	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected table name")
	}
	stmt.TableName = p.curTok.Literal
	p.nextToken()

	// Parse optional WHERE clause
	if p.curTok.Type == TokenWhere {
		where, err := p.parseWhere()
		if err != nil {
			return nil, err
		}
		stmt.Where = where
	}

	return stmt, nil
}

// parseCreate parses a CREATE TABLE statement
func (p *Parser) parseCreate() (*CreateTableStatement, error) {
	stmt := &CreateTableStatement{}

	// Parse TABLE keyword
	p.nextToken()
	if err := p.expectToken(TokenTable); err != nil {
		return nil, err
	}
	p.nextToken()

	// Parse table name
	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected table name")
	}
	stmt.TableName = p.curTok.Literal
	p.nextToken()

	// Parse column definitions (stub)
	stmt.Columns = []ColumnDef{
		{Name: "id", Type: "int"},
	}

	return stmt, nil
}

// parseDrop parses a DROP TABLE statement
func (p *Parser) parseDrop() (*DropTableStatement, error) {
	stmt := &DropTableStatement{}

	// Parse TABLE keyword
	p.nextToken()
	if err := p.expectToken(TokenTable); err != nil {
		return nil, err
	}
	p.nextToken()

	// Parse table name
	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected table name")
	}
	stmt.TableName = p.curTok.Literal

	return stmt, nil
}

// parseWhere parses a WHERE clause
func (p *Parser) parseWhere() (*WhereClause, error) {
	where := &WhereClause{
		Conditions: []Condition{},
	}

	p.nextToken()

	// Parse conditions (simplified - single condition)
	if p.curTok.Type != TokenIdentifier {
		return nil, fmt.Errorf("expected column name in WHERE")
	}
	col := p.curTok.Literal
	p.nextToken()

	var op string
	switch p.curTok.Type {
	case TokenEqual:
		op = "="
	case TokenNotEqual:
		op = "<>"
	case TokenLessThan:
		op = "<"
	case TokenGreaterThan:
		op = ">"
	default:
		return nil, fmt.Errorf("expected operator in WHERE")
	}
	p.nextToken()

	var val interface{}
	switch p.curTok.Type {
	case TokenNumber:
		num, _ := strconv.Atoi(p.curTok.Literal)
		val = num
	case TokenString:
		val = p.curTok.Literal
	default:
		return nil, fmt.Errorf("expected value in WHERE")
	}
	p.nextToken()

	where.Conditions = append(where.Conditions, Condition{
		Column:   col,
		Operator: op,
		Value:    val,
	})

	return where, nil
}
