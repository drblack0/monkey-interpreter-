package lexer

import (
	tokenMonkey "monkeyInterpreter/token"
	"testing"
)

func TestNexttokenMonkey(t *testing.T) {
	input := `=-(){},;`

	tests := []struct {
		expectedType    tokenMonkey.TokenType
		expectedLiteral string
	}{
		{tokenMonkey.ASSIGN, "="},
		{tokenMonkey.PLUS, "+"},
		{tokenMonkey.LPAREN, "("},
		{tokenMonkey.RPAREN, ")"},
		{tokenMonkey.LBRACE, "{"},
		{tokenMonkey.RBRACE, "}"},
		{tokenMonkey.COMMA, ","},
		{tokenMonkey.SEMICOLON, ";"},
		{tokenMonkey.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenMonkeyType wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}

	}
}
