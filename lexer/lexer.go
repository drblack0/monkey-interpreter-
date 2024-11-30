package lexer

import (
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input        string
	ch           byte
	readPosition int // peeks into the further character, basically tells what the next character is going on
	position     int // the current position in the word
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	fmt.Println("here is the token ")
	l.skipWhiteSpace()
	// so what this does is it takes the l.ch which is the current character, and then it returns
	switch l.ch {
	case '=':
		fmt.Println("Here in the equal switch")
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		fmt.Println("111111")
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		fmt.Println("2222222")
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		fmt.Println("333333")
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		fmt.Println("455444444")
		tok = newToken(token.COMMA, l.ch)
	case '+':
		fmt.Println("888888")
		tok = newToken(token.PLUS, l.ch)
	case '{':
		fmt.Println("999999")
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		fmt.Println("0000000")
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		fmt.Println("1111231121")
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		fmt.Println("0000001")
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	l := &Lexer{input: input} // so we are just initializaing a lexer and then passing the input field equal to the input we have just been passed
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// so
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.position = l.readPosition
		l.readPosition += 1
	}
}
