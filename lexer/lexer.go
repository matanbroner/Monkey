package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input          string
	position       int
	readAtPosition int
	ch             byte
}

func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar() // init our reader
	return lex
}

func (lex *Lexer) readChar() {
	if lex.readAtPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readAtPosition]
	}
	lex.position = lex.readAtPosition
	lex.readAtPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tok *token.Token
	switch lex.ch {
	case '=':
		tok = token.New(token.ASSIGN, lex.ch)
	case '+':
		tok = token.New(token.PLUS, lex.ch)
	case '-':
		tok = token.New(token.MINUS, lex.ch)
	case ';':
		tok = token.New(token.SEMICOLON, lex.ch)
	case '{':
		tok = token.New(token.LBRACE, lex.ch)
	case '}':
		tok = token.New(token.RBRACE, lex.ch)
	case '(':
		tok = token.New(token.LPAREN, lex.ch)
	case ')':
		tok = token.New(token.RPAREN, lex.ch)
	case 0:
		tok = token.New(token.EOF, 0)
	default:
		tok = token.New(token.ILLEGAL, 0) // TODO: replace with char for illegal
	}
	lex.readChar()
	return *tok
}
