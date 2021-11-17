// Package Lexer provides with data structures and methods for tokenization monkeylang soruce code.
package lexer

import "github.com/kellemNegasi/monkeylang/token"

// Lexer defiens a new struct type that represents the Lexer
type Lexer struct {
	input        string // the soruce code input.
	position     int    // current position in input (current char)
	readPosition int    // the position after current char, current reading position
	ch           byte   // current char under examination
}

// New(): - A function that intializes Lexer object and returns a pointer to a it.
// input: string - the raw soource code fed to the lexer.
func new(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}
	l.readChar() // initializes position,readPosition and ch.
	return l
}

// readChar() Reads the next character and assigns it the ch field of Lexer.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken(): identifies and returns the next token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	}
	l.readChar()
	return tok
}

// newToken initializes new Token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
