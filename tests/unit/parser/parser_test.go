package parser

import (
	"testing"

	"github.com/Gi-v/db-flow-AutomateSchemaEvolutionPipeline/src/parser"
	"github.com/stretchr/testify/assert"
)

func TestLexerBasic(t *testing.T) {
	input := "SELECT * FROM users"
	lexer := parser.NewLexer(input)

	tok := lexer.NextToken()
	assert.Equal(t, parser.TokenSelect, tok.Type)

	tok = lexer.NextToken()
	assert.Equal(t, parser.TokenStar, tok.Type)

	tok = lexer.NextToken()
	assert.Equal(t, parser.TokenFrom, tok.Type)

	tok = lexer.NextToken()
	assert.Equal(t, parser.TokenIdentifier, tok.Type)
	assert.Equal(t, "users", tok.Literal)
}

func TestParserSelect(t *testing.T) {
	p := parser.NewParser()
	
	ast, err := p.Parse("SELECT * FROM users")
	assert.NoError(t, err)
	assert.NotNil(t, ast)

	selectStmt, ok := ast.(*parser.SelectStatement)
	assert.True(t, ok)
	assert.Equal(t, "users", selectStmt.TableName)
	assert.Contains(t, selectStmt.Columns, "*")
}

func TestParserInsert(t *testing.T) {
	p := parser.NewParser()
	
	ast, err := p.Parse("INSERT INTO users VALUES (1, 'John')")
	assert.NoError(t, err)
	assert.NotNil(t, ast)

	insertStmt, ok := ast.(*parser.InsertStatement)
	assert.True(t, ok)
	assert.Equal(t, "users", insertStmt.TableName)
}

func TestParserUpdate(t *testing.T) {
	p := parser.NewParser()
	
	ast, err := p.Parse("UPDATE users SET name = 'Jane' WHERE id = 1")
	assert.NoError(t, err)
	assert.NotNil(t, ast)

	updateStmt, ok := ast.(*parser.UpdateStatement)
	assert.True(t, ok)
	assert.Equal(t, "users", updateStmt.TableName)
}

func TestParserDelete(t *testing.T) {
	p := parser.NewParser()
	
	ast, err := p.Parse("DELETE FROM users WHERE id = 1")
	assert.NoError(t, err)
	assert.NotNil(t, ast)

	deleteStmt, ok := ast.(*parser.DeleteStatement)
	assert.True(t, ok)
	assert.Equal(t, "users", deleteStmt.TableName)
}
