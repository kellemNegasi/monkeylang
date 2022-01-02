// Package ast defines functions and data structures to build an Abstract Syntax Tree for parsing a source code.
package ast

import "github.com/kellemNegasi/monkeylang/token"

// Node defines the basic node interface that represents any kind of node in the AST.
type Node interface {
	TokenLiteral() string
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
	Value *Expression
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
