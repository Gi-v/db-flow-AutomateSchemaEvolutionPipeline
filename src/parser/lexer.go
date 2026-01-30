package parser

import (
	"fmt"
	"unicode"
)

// TokenType represents the type of a token
type TokenType int

const (
	// Keywords
	TokenSelect TokenType = iota
	TokenInsert
	TokenUpdate
	TokenDelete
	TokenFrom
	TokenWhere
	TokenInto
	TokenValues
	TokenSet
	TokenCreate
	TokenTable
	TokenDrop

	// Operators
	TokenEqual
	TokenNotEqual
	TokenLessThan
	TokenGreaterThan
	TokenAnd
	TokenOr

	// Literals
	TokenIdentifier
	TokenNumber
	TokenString

	// Punctuation
	TokenComma
	TokenSemicolon
	TokenLeftParen
	TokenRightParen
	TokenStar

	// Special
	TokenEOF
	TokenIllegal
)

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Lexer performs lexical analysis on SQL input
type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
	line    int
	column  int
}

// NewLexer creates a new lexer for the given input
func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

// readChar reads the next character and advances position
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
	l.column++

	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
}

// peekChar looks at the next character without advancing
func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	tok.Line = l.line
	tok.Column = l.column

	switch l.ch {
	case '=':
		tok = Token{Type: TokenEqual, Literal: "="}
	case ',':
		tok = Token{Type: TokenComma, Literal: ","}
	case ';':
		tok = Token{Type: TokenSemicolon, Literal: ";"}
	case '(':
		tok = Token{Type: TokenLeftParen, Literal: "("}
	case ')':
		tok = Token{Type: TokenRightParen, Literal: ")"}
	case '*':
		tok = Token{Type: TokenStar, Literal: "*"}
	case '<':
		if l.peekChar() == '>' {
			l.readChar()
			tok = Token{Type: TokenNotEqual, Literal: "<>"}
		} else {
			tok = Token{Type: TokenLessThan, Literal: "<"}
		}
	case '>':
		tok = Token{Type: TokenGreaterThan, Literal: ">"}
	case '\'', '"':
		tok.Type = TokenString
		tok.Literal = l.readString(l.ch)
	case 0:
		tok = Token{Type: TokenEOF, Literal: ""}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupKeyword(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = TokenNumber
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = Token{Type: TokenIllegal, Literal: string(l.ch)}
		}
	}

	l.readChar()
	return tok
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readIdentifier reads an identifier or keyword
func (l *Lexer) readIdentifier() string {
	pos := l.pos
	for isLetter(l.ch) || isDigit(l.ch) || l.ch == '_' {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

// readNumber reads a number
func (l *Lexer) readNumber() string {
	pos := l.pos
	for isDigit(l.ch) || l.ch == '.' {
		l.readChar()
	}
	return l.input[pos:l.pos]
}

// readString reads a string literal
func (l *Lexer) readString(quote byte) string {
	l.readChar() // consume opening quote
	pos := l.pos
	for l.ch != quote && l.ch != 0 {
		l.readChar()
	}
	str := l.input[pos:l.pos]
	return str
}

// isLetter checks if a character is a letter
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch))
}

// isDigit checks if a character is a digit
func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

// lookupKeyword maps string to keyword token type
func lookupKeyword(ident string) TokenType {
	keywords := map[string]TokenType{
		"select": TokenSelect,
		"SELECT": TokenSelect,
		"insert": TokenInsert,
		"INSERT": TokenInsert,
		"update": TokenUpdate,
		"UPDATE": TokenUpdate,
		"delete": TokenDelete,
		"DELETE": TokenDelete,
		"from":   TokenFrom,
		"FROM":   TokenFrom,
		"where":  TokenWhere,
		"WHERE":  TokenWhere,
		"into":   TokenInto,
		"INTO":   TokenInto,
		"values": TokenValues,
		"VALUES": TokenValues,
		"set":    TokenSet,
		"SET":    TokenSet,
		"create": TokenCreate,
		"CREATE": TokenCreate,
		"table":  TokenTable,
		"TABLE":  TokenTable,
		"drop":   TokenDrop,
		"DROP":   TokenDrop,
		"and":    TokenAnd,
		"AND":    TokenAnd,
		"or":     TokenOr,
		"OR":     TokenOr,
	}

	if tokType, ok := keywords[ident]; ok {
		return tokType
	}
	return TokenIdentifier
}

// TokenString returns a string representation of token type
func (tt TokenType) String() string {
	names := map[TokenType]string{
		TokenSelect:      "SELECT",
		TokenInsert:      "INSERT",
		TokenUpdate:      "UPDATE",
		TokenDelete:      "DELETE",
		TokenFrom:        "FROM",
		TokenWhere:       "WHERE",
		TokenInto:        "INTO",
		TokenValues:      "VALUES",
		TokenSet:         "SET",
		TokenCreate:      "CREATE",
		TokenTable:       "TABLE",
		TokenDrop:        "DROP",
		TokenEqual:       "=",
		TokenNotEqual:    "<>",
		TokenLessThan:    "<",
		TokenGreaterThan: ">",
		TokenAnd:         "AND",
		TokenOr:          "OR",
		TokenIdentifier:  "IDENTIFIER",
		TokenNumber:      "NUMBER",
		TokenString:      "STRING",
		TokenComma:       ",",
		TokenSemicolon:   ";",
		TokenLeftParen:   "(",
		TokenRightParen:  ")",
		TokenStar:        "*",
		TokenEOF:         "EOF",
		TokenIllegal:     "ILLEGAL",
	}
	if name, ok := names[tt]; ok {
		return name
	}
	return fmt.Sprintf("Unknown(%d)", tt)
}
