package lexer

import (
	tokenMonkey "monkeyInterpreter/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() tokenMonkey.Token {
	var tok tokenMonkey.Token

	switch l.ch {
	case '=':
		tok = newToken(tokenMonkey.ASSIGN, l.ch)
	case ';':
		tok = newToken(tokenMonkey.SEMICOLON, l.ch)
	case '(':
		tok = newToken(tokenMonkey.LPAREN, l.ch)
	case ')':
		tok = newToken(tokenMonkey.RPAREN, l.ch)
	case ',':
		tok = newToken(tokenMonkey.COMMA, l.ch)
	case '+':
		tok = newToken(tokenMonkey.PLUS, l.ch)
	case '{':
		tok = newToken(tokenMonkey.ASSIGN, l.ch)
	case '}':
		tok = newToken(tokenMonkey.ASSIGN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tokenMonkey.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType tokenMonkey.TokenType, ch byte) tokenMonkey.Token {
	return tokenMonkey.Token{
		Type: tokenType, Literal: string(ch),
	}
}
