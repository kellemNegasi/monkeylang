// Package token provids with types and methods for dealing with and manipulating tokens.
package token

// TokenType defines a string Type that represents the type of a given token
type TokenType string

// Token is a struct that represents a single token object
type Token struct {
	Type    TokenType
	Literal string
}

// keywords defiens map of keywords in the language.

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

// LookupIdent checks a given keyword wether it is an identifier or a keyword.
// First it checks the keywords tabble and returns the type of keyword if found, otherwise it returns IDENT.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Defenitions of Token names.
const (
	// ILLEGAL defines an illegal token.
	ILLEGAL = "ILLEGAL"
	// EOF defines an end of file token.
	EOF = "EOF"
	// IDENT and INT identifiers and literals.
	IDENT = "IDENT" // variables and function names
	// INT token represents integre variables i.e 123456789.
	INT = "INT"

	// ASSIGN and other operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOTEQ    = "!="
	// COMMA and other delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// FUNCTION and other keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)
