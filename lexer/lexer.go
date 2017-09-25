package lexer

import "github.com/kscarlett/monkey/token"

type Lexer struct {
	input        string
	position     int  // position is the current position in input (points to the current char)
	readPosition int  // readPosition is the current reading position in input (after current char)
	ch           byte // ch is the current char under examination — change to rune and read in differently for Unicode support (pg.15)
}

func New(input string) *Lexer {
	/*
	 * Improvement: initialise with an io.reader and the filename
	 * instead of a string to attach filenames and line numbers to the tokens,
	 * to better track down lexting and parsing errors (pg.13/14)
	 */
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
