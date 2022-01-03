package parser

import (
	"fmt"

	"github.com/kellemNegasi/monkeylang/ast"
	"github.com/kellemNegasi/monkeylang/lexer"
	"github.com/kellemNegasi/monkeylang/token"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

// Parser reprsents the parser object.
type Parser struct {
	lexer          *lexer.Lexer
	currentToken   token.Token
	peekToken      token.Token
	errors         []string // for holding the errors.
	prefixParseFns map[token.TokenType]prefixParserFn
	infixParseFns  map[token.TokenType]infixParserFn
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
	// associate tokens with corrosponding parsing functions.
	p.prefixParseFns = make(map[token.TokenType]prefixParserFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	return p
}

// parseIdentifier parses an indentifier and returns Identifier object.

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
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
		return p.ParseLetStatement()
	case token.RETURN:
		return p.ParseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// ParseLetStatment is a specific statment parser that is dedicated to parsing a `let` statment.
func (p *Parser) ParseLetStatement() *ast.LetStatement {
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
func (p *Parser) ParseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.currentToken}
	p.nextToken()

	// TODO: We're skipping the expression parsing

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

type prefixParserFn func() ast.Expression

type infixParserFn func(ast.Expression) ast.Expression

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParserFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParserFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.ParseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

// parseEXpression parses a given expression and returns the left expression given an operator.
func (p *Parser) ParseExpression(precedense int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}
