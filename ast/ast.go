package ast

import (
	"bytes"

	"github.com/esquivias/interpreter/token"
)

// Node interface implemented by every node in our AST
type Node interface {
	// TokenLiteral() will only be used for debugging and testing
	TokenLiteral() string
	String() string
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

// TokenLiteral function on Identifier struct (Expression interface)
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// expressionNode function on Identifier struct (Expression interface)
func (i *Identifier) expressionNode() {}

func (i *Identifier) String() string { return i.Value }

/*
 * Program
 */

// Program struct contains an array of Statement Nodes
type Program struct {
	Statements []Statement
}

// TokenLiteral function on Program struct
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

/*
 * Integer Literal
 */

// IntegerLiteral struct
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// expressionNode function on IntegerLiteral
func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral function on IntegerLiteral
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String function on IntegerLiteral
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

/*
 * LetStatement
 */

// LetStatement struct
type LetStatement struct {
	// let x = 5
	Token token.Token // token (token.LET)
	Name  *Identifier // identifier of the binding (token.IDENT, x)
	Value Expression  // expression that produces the value (INT 5)
}

// statementNode function on LetStatement
func (ls *LetStatement) statementNode() {}

// TokenLiteral function on LetStatement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

/*
 * Expression Statement
 */

// ExpressionStatement struct
type ExpressionStatement struct {
	Token      token.Token // the first token in the expression
	Expression Expression
}

// statementNode function on ExpressionStatement
func (es *ExpressionStatement) statementNode() {

}

// TokenLiteral function on ExpressionStatement
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

/*
 * Return Statement
 */

// ReturnStatement struct consist soley of the keyword return and an expression
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

// statementNode function on ReturnStatement
func (rs *ReturnStatement) statementNode() {}

// TokenLiteral function on ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// PrefixExpression struct
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// expressionNode function on PrefixExpression struct
func (pe *PrefixExpression) expressionNode() {

}

// TokenLiteral function
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String function on PrefixExpression struct
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression struct
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

// expressionNode function on InfixExpression
func (ie *InfixExpression) expressionNode() {

}

// TokenLiteral function on InfixExpression
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

// String function on InfixExpression
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}
