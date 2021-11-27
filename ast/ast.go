// Package ast defines functions and data structures to build an Abstract Syntax Tree for parsing a source code.
package ast

// Node defines the basic node interface that represents any kind of node in the AST.
type Node interface {
	TokenLiteral() string
}

// Statement defines the Statment node interface.
type Statement interface {
	Node
	statmentNode()
}

// Expression defines an Expression node.
type Expression interface {
	Node
	ExpressionNode()
}

// Program defines a type that represents source code program which is made up of one or more statments.
type Program struct {
	Statement []Statement
}

// TokenLiteral is a method that returns the litral value of the token the node is associated with.
func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	}

	return ""
}
