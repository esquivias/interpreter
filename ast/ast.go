package ast

import (
	"github.com/esquivias/interpreter/token"
)

// Node interface is implemented by every node in our AST
type Node interface {
	// TokenLiteral() will only be used for debugging and testing
	TokenLiteral() string
}

// Statement interface implements Node, must provide a TokenLiteral() method
type Statement interface {
	Node
	statementNode()
}

// Expression interface implements Node, must provide a TokenLiteral() method
type Expression interface {
	Node
	expressionNode()
}

// Identifier struct implements the Expression interface in case a value is produced
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral method for Identifier struct (Expression interface)
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// expressionNode method for Identifier struct (Expression interface)
func (i *Identifier) expressionNode() {}

// Program struct contains an array of Statement Nodes
type Program struct {
	Statements []Statement
}

// TokenLiteral method for Program struct (Statement interface (Node interface))
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

//

// LetStatement has two methods, statementNode and TokenLiteral, to satisfy the Statement and Node interfaces respectively.
type LetStatement struct {
	// let x = 5
	Token token.Token // token (token.LET)
	Name  *Identifier // identifier of the binding (token.IDENT, x)
	Value Expression  // expression that produces the value (INT 5)
}

// statementNode method for LetStatement.Value.Expression struct
func (ls *LetStatement) statementNode() {}

// TokenLiteral method for LetStatement.Value.Expression struct
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
