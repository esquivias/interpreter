package parser

import (
	"fmt"

	"github.com/esquivias/interpreter/ast"
	"github.com/esquivias/interpreter/lexer"
	"github.com/esquivias/interpreter/token"
)

// Parser struct
type Parser struct {
	l         *lexer.Lexer // pointer to an instance of the lexer (NextToken())
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

// New Parser returns a Parser struct with a lexer and tokens set.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// Read two tokens so both curToken and peekToken are set
	p.nextToken()
	p.nextToken()
	return p
}

// ParseProgram method for Parser struct returns AST Program Node
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{} // construct the root node of the AST
	program.Statements = []ast.Statement{}
	// iterate over every token in the input until an token.EOF token is encountered
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// parseStatement method based on defined token types
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement returns a LET Statement AST Node
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// TODO: We're skipping the expressions until we encounter a semicolon
	// NOTE: A missing semicolon will cause a timeout
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// nextToken method sets the parser's current token and peek token
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenIs returns true if the parser's current token type is the provided token type
func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

// peekTokenIs returns true if the parser's peek token type is the provided token type
func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

// expectPeek method advances the current and peek token and returns true if the parser's peek token type is the provided token type
func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// Errors returns parser errors array
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError appends an error message to the parser errors array
func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
