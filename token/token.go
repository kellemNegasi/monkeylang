package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

// keywords defiens map of keywords in the language.

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent() checks a given keyword wether it is an identifier or a keyword.
// First it checks the keywords tabble and returns the type of keyword if found, otherwise it returns IDENT.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
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
