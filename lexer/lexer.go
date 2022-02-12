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
	// consume next char
	if lex.readAtPosition >= len(lex.input) {
		lex.ch = 0
	} else {
		lex.ch = lex.input[lex.readAtPosition]
	}
	lex.position = lex.readAtPosition
	lex.readAtPosition += 1
}

func (lex *Lexer) peekChar() byte {
	// peek but do not consume
	if lex.readAtPosition >= len(lex.input) {
		return 0
	} else {
		return lex.input[lex.readAtPosition]
	}
}

func (lex *Lexer) NextToken() token.Token {
	lex.skipWhitespaces() // prevent whitespaces in tokens
	var tok *token.Token
	switch lex.ch {
	case '=':
		if lex.peekChar() == '=' {
			lex.readChar()
			tok = &token.Token{Type: token.EQ, Literal: "=="}
		} else {
			tok = token.New(token.ASSIGN, lex.ch)
		}
	case '!':
		if lex.peekChar() == '=' {
			lex.readChar()
			tok = &token.Token{Type: token.NEQ, Literal: "!="}
		} else {
			tok = token.New(token.BANG, lex.ch)
		}
	case '+':
		tok = token.New(token.PLUS, lex.ch)
	case '-':
		tok = token.New(token.MINUS, lex.ch)
	case '*':
		tok = token.New(token.ASTERISK, lex.ch)
	case '/':
		tok = token.New(token.SLASH, lex.ch)
	case '<':
		tok = token.New(token.LT, lex.ch)
	case '>':
		tok = token.New(token.GT, lex.ch)
	case ';':
		tok = token.New(token.SEMICOLON, lex.ch)
	case ',':
		tok = token.New(token.COMMA, lex.ch)
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
		tok = token.New(token.EOF, 0) // create dummy token
		if isLetter(lex.ch) {         // find identifier or keyword
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupKeyword(tok.Literal)
			return *tok // we already rea through the end, no need for lex.readChar()
		} else if isDigit(lex.ch) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return *tok
		} else {
			tok.Type = token.ILLEGAL
			tok.Literal = string(lex.ch)
		}
	}
	lex.readChar()
	return *tok
}

func (lex *Lexer) readIdentifier() string {
	position := lex.position
	for isLetter(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func (lex *Lexer) readNumber() string {
	position := lex.position
	for isDigit(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lex *Lexer) skipWhitespaces() {
	for lex.ch == ' ' || lex.ch == '\t' || lex.ch == '\r' || lex.ch == '\n' {
		lex.readChar()
	}
}
