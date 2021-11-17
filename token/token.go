package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// identifiers and literals
	IDENT = "IDENT" // variables and function names
	INT   = "INT"   // integers (123456789)

	// operators
	ASSIGN = "=" //
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
