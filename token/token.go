package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // x, y, add
	INT   = "INT"

	// Ops
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// Syntax Chars
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

func New(tokenType TokenType, ch byte) *Token {
	var literal string
	if ch == 0 {
		literal = ""
	} else {
		literal = string(ch)
	}
	tok := &Token{Type: tokenType, Literal: literal}
	return tok
}
