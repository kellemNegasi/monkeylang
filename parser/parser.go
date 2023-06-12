package parser

import (
	"fmt"
	"strconv"

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
	PREFIX      // -X or !XCALL
)

var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOTEQ:    EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

// Parser reprsents the parser object.
type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
	errors       []string // for holding the errors.

	// maps associating infix and prefix operator tokens to appropriate parser functions
	infixParseFns  map[token.TokenType]infixParseFn
	prefixParseFns map[token.TokenType]prefixParseFn
}

// New initializes a Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefixParser(token.IDENT, p.parseIdentifier)
	p.registerPrefixParser(token.INT, p.parseIntegerLiteral)
	p.registerPrefixParser(token.BANG, p.parsePrefixExpression)
	p.registerPrefixParser(token.MINUS, p.parsePrefixExpression)

	// register infix parsing functions

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfixParser(token.PLUS, p.parseInfixExpression)
	p.registerInfixParser(token.MINUS, p.parseInfixExpression)
	p.registerInfixParser(token.SLASH, p.parseInfixExpression)
	p.registerInfixParser(token.ASTERISK, p.parseInfixExpression)
	p.registerInfixParser(token.EQ, p.parseInfixExpression)
	p.registerInfixParser(token.NOTEQ, p.parseInfixExpression)
	p.registerInfixParser(token.LT, p.parseInfixExpression)
	p.registerInfixParser(token.GT, p.parseInfixExpression)

	// warm start the parser with two tokens i.e one for currentToken and the next for peekToken.
	p.nextToken()
	p.nextToken()
	return p
}

// Function types for parsing prefix and infix operation expression.
type (
	infixParseFn  func(ast.Expression) ast.Expression
	prefixParseFn func() ast.Expression
)

// registerPrefixParser adds a parser for a given prefix operator token t.
func (p *Parser) registerPrefixParser(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

// registerInfixParser adds an infix parsing function for the given infix operator token t.
func (p *Parser) registerInfixParser(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
}

func (p *Parser) ParseExpressionStatment() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	leftExp := prefix()
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	val, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = val
	return lit
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
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
		return p.ParseExpressionStatment()
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
	// If not a semicolon get the next token.
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}
	return LOWEST
}
