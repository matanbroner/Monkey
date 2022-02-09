package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "+{}-="
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.PLUS, "+"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.MINUS, "-"},
		{token.ASSIGN, "="},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("TestNextToken[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("TestNextToken[%d] - token literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
