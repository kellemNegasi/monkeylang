package parser

import (
	"fmt"

	"github.com/kellemNegasi/monkeylang/ast"
	"github.com/kellemNegasi/monkeylang/lexer"
	"github.com/kellemNegasi/monkeylang/token"
)

// Parser reprsents the parser object.
type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string // for holding the errors.
}

// New initializes a Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}
	// warm start the parser with two tokens i.e one for currentToken and the next for peekToken.
	p.nextToken()
	p.nextToken()
	return p
}

// Errors returns the parser errors.
func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken gets the imediate next token in program.
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

// ParseProgram parses the program
func (p *Parser) ParseProgram() *ast.Program {
	program := ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currentToken.Type != token.EOF {
		statement := p.ParseStatment()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return &program
}

// ParseStatment parses a given statemnt.
func (p *Parser) ParseStatment() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.ParseLetStatment()
	default:
		return nil
	}
}

// ParseLetStatment is a specific statment parser that is dedicated to parsing a `let` statment.
func (p *Parser) ParseLetStatment() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}
	// after the `let` keyword check the next token's type is an Identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	// if identifier is found in the next token the construct an Identifier object.
	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}
	// If the next token is not a semicolon parse the experssion.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

// curTokenIs checks the currentToken is the same as `t`.
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

// peekTokenI checks if the token type of the peakToken is the same as `t.TokenType`.
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek checks the if the peakToken is equal to `t` and advances the parser to next Token.
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

// peekError adds an error to errors when the type of peekToken doesn’t match the expectation.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
