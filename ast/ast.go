// Package ast defines functions and data structures to build an Abstract Syntax Tree for parsing a source code.
package ast

import (
	"bytes"

	"github.com/kellemNegasi/monkeylang/token"
)

// Node defines the basic node interface that represents any kind of node in the AST.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement defines the Statement node interface.
type Statement interface {
	Node
	statementNode()
}

// Expression defines an Expression node.
type Expression interface {
	Node
	ExpressionNode()
}

// Program defines a type that represents source code program which is made up of one or more statements.
type Program struct {
	Statements []Statement
}

// TokenLiteral is a method that returns the litral value of the token the node is associated with.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// Identifier represents identifier type.
type Identifier struct {
	Token token.Token
	Value string
}

// LetStatement represents a node for let statement binding.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (letStmnt *LetStatement) statementNode() {
}

//TokenLiteral returns the literal value of the Token field of the LetStatement.
func (letStmnt *LetStatement) TokenLiteral() string {
	return letStmnt.Token.Literal
}

// ExpressionNode makes Identifier implement the Expression interface.
func (id *Identifier) ExpressionNode() {

}

// TokenLiteral method of Identifier makes it implement the Node interface.
func (id *Identifier) TokenLiteral() string {
	return id.Token.Literal
}

type ReturnStatement struct {
	Token       token.Token // the `return` token is represented here.
	ReturnValue Expression
}

// statementNode implements the Statement interface
func (retStatement *ReturnStatement) statementNode() {}

// TokenLiteral returns the literal value of retStatement.Token attribute.
func (retStatement *ReturnStatement) TokenLiteral() string {
	return retStatement.Token.Literal
}

type ExpressionStatement struct {
	Token      token.Token // this holds the first tokne of the expression
	Expression Expression
}

// statementNode implements the Statement interface.
func (exp *ExpressionStatement) statementNode() {
}

func (expStmt *ExpressionStatement) TokenLiteral() string {
	return expStmt.Token.Literal
}

// String implements the Node interface

func (i *Identifier) String() string { return i.Value }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
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
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) ExpressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pexp *PrefixExpression) ExpressionNode()      {}
func (pexp *PrefixExpression) TokenLiteral() string { return pexp.Token.Literal }
func (pexp *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pexp.Operator)
	out.WriteString(pexp.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) ExpressionNode() {

}
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")
	return out.String()
}
