package lexer

import (
	"testing"

	"github.com/kellemNegasi/monkeylang/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLIteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}
	l := new(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d]-- tokentype wrong. expected=%q,got =%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLIteral {
			t.Fatalf("tests[%d] --literal wrong. expected %q, got %q", i, tt.expectedLIteral, tok.Literal)
		}
	}
}
