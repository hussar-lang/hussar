package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"

	"hussar.io/lang/token"
)

// TODO: finish this refactoring with this: https://blog.gopheracademy.com/advent-2014/parsers-lexers/

type Lexer struct {
	input     *bufio.Reader
	currentCh rune
	peekCh    rune
	lastToken token.Token

	currentFile string
	line        int
	column      int
}

func New(reader io.Reader) *Lexer {
	l := &Lexer{input: bufio.NewReader(reader)}

	// Initialise current and peeked characters
	l.readRune()
	l.readRune()
	l.column = 1
	l.line = 1
	l.currentFile = "buffer"
	return l
}

func NewFile(file string) (*Lexer, error) {
	l := &Lexer{currentFile: file}
	if err := l.loadFile(); err != nil {
		return nil, err
	}

	// Initialise current and peeked characters
	l.readRune()
	l.readRune()
	l.column = 1
	l.line = 1
	return l, nil
}

func NewString(input string) *Lexer {
	return New(strings.NewReader(input))
}

func (l *Lexer) loadFile() error {
	if l.currentFile == "" {
		panic("No files to load")
	}

	file, err := os.Open(l.currentFile)
	if err != nil {
		return fmt.Errorf("Failed to open file %s", l.currentFile)
	}

	l.input = bufio.NewReader(file)
	return nil
}

func (l *Lexer) readRune() {
	oldPeek := l.peekCh

	newPeek, _, err := l.input.ReadRune()
	if err != nil {
		if l.currentFile != "" {
			if err := l.loadFile(); err != nil {
				panic(err)
			}
			l.readRune()
			return
		}
		l.peekCh = 0
		l.currentCh = oldPeek
		return
	}
	l.currentCh = oldPeek
	l.peekCh = newPeek
	l.column++
}

func makePos(line, column int) token.Position {
	return token.Position{
		Line:   line,
		Column: column,
	}
}

func (l *Lexer) curPosition() token.Position {
	return makePos(l.line, l.column)
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.trimWhitespaceNotNewLine()

	switch l.currentCh {
	case '\n':
		if l.needSemicolon() {
			tok = l.newToken(token.SEMICOLON, ';')
			l.resetPos()
		} else {
			l.trimWhitespace()
			return l.NextToken()
		}
		// case '=':
		// 	if l.peekCharIs('=') {
		// 		ch := l.ch
		// 		l.readChar()
		// 		literal := string(ch) + string(l.ch)
		// 		tok = token.Token{Type: token.EQ, Literal: literal}
		// 	} else {
		// 		tok = newToken(token.ASSIGN, l.ch)
		// 	}
		// case '+':
		// 	tok = newToken(token.PLUS, l.ch)
		// case '-':
		// 	tok = newToken(token.MINUS, l.ch)
		// case '!':
		// 	if l.peekCharIs('=') {
		// 		ch := l.ch
		// 		l.readChar()
		// 		literal := string(ch) + string(l.ch)
		// 		tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		// 	} else {
		// 		tok = newToken(token.BANG, l.ch)
		// 	}
		// case '/':
		// 	tok = newToken(token.SLASH, l.ch)
		// case '*':
		// 	tok = newToken(token.ASTERISK, l.ch)
		// case '<':
		// 	tok = newToken(token.LT, l.ch)
		// case '>':
		// 	tok = newToken(token.GT, l.ch)
		// case ';':
		// 	tok = newToken(token.SEMICOLON, l.ch)
		// case ',':
		// 	tok = newToken(token.COMMA, l.ch)
		// case '(':
		// 	tok = newToken(token.LPAREN, l.ch)
		// case ')':
		// 	tok = newToken(token.RPAREN, l.ch)
		// case '{':
		// 	tok = newToken(token.LBRACE, l.ch)
		// case '}':
		// 	tok = newToken(token.RBRACE, l.ch)
		// case '[':
		// 	tok = newToken(token.LBRACKET, l.ch)
		// case ']':
		// 	tok = newToken(token.RBRACKET, l.ch)
		// case '"':
		// 	tok.Type = token.STRING
		// 	tok.Literal = l.readString()
		// case 0:
		// 	tok.Literal = ""
		// 	tok.Type = token.EOF
		// default:
		// 	if isLetter(l.ch) {
		// 		tok.Literal = l.readIdentifier()
		// 		tok.Type = token.LookupIdent(tok.Literal)
		// 		return tok
		// 	} else if isDigit(l.ch) {
		// 		tok.Type = token.INT // Improvement: To be changed when adding more numeric types (float, hex, oct, binary)
		// 		tok.Literal = l.readNumber()
		// 		return tok
		// 	} else {
		// 		tok = newToken(token.ILLEGAL, l.ch)
		// 	}
	}

	// l.readChar()
	// return tok
}

func (l *Lexer) peekChar() rune {
	return l.peekCh
}

func (l *Lexer) resetPos() {
	l.line++
	l.column = 0
}

func (l *Lexer) needSemicolon() bool {
	return l.lastTokenWas(
		token.IDENT,
		token.INT,
		token.STRING,
		token.NIL,
		token.RETURN,
		token.RPAREN,
		token.RBRACKET,
		token.RBRACE) && !l.lastTokenWas(token.SEMICOLON)
}

func (l *Lexer) readIdentifier() string {
	var ident bytes.Buffer

	for isIdent(l.currentCh) {
		ident.WriteRune(l.currentCh)
		l.readRune()
	}

	return ident.String()
}

func (l *Lexer) readString() token.Token {
	var ident bytes.Buffer
	pos := l.curPosition()
	l.readRune()

	for l.currentCh != '"' {
		if l.currentCh == '\n' {
			return token.Token{
				Literal:  "Newlines are not allowed in strings",
				Type:     token.ILLEGAL,
				Pos:      l.curPosition(),
				Filename: l.currentFile,
			}
		}

		if l.currentCh == '\\' {
			l.readRune()

			switch l.currentCh {
			case '0': // null
				ident.WriteRune(0)
			case 'b': // backspace
				ident.WriteRune('\b')
			case 'n': // newline
				ident.WriteRune('\n')
			case 'r': // carriage return
				ident.WriteRune('\r')
			case 't': // horizontal tab
				ident.WriteRune('\t')
			case 'v': // vertical tab
				ident.WriteRune('\v')
			case 'f': // form feed
				ident.WriteRune('\f')
			case 'e': // escape
				ident.WriteRune('\033')
			case '\\': // backslash
				ident.WriteRune('\\')
			case '"': // double quote
				ident.WriteRune('"')
			default:
				ident.WriteByte('\\')
				ident.WriteRune(l.currentCh)
			}
			l.readRune()
			continue
		}
		ident.WriteRune(l.currentCh)
		l.readRune()
	}

	return token.Token{
		Literal:  ident.String(),
		Type:     token.STRING,
		Pos:      pos,
		Filename: l.currentFile,
	}
}

func (l *Lexer) readRawString() token.Token {
	var ident bytes.Buffer
	pos := l.curPosition()
	l.readRune()

	for l.currentCh != '\'' {
		if l.currentCh == '\\' && l.peekCh == '\'' {
			l.readRune() // Go past the backslash so the next line will be a single quote
		}
		ident.WriteRune(l.currentCh)
		l.readRune()
	}

	return token.Token{
		Literal:  ident.String(),
		Type:     token.STRING,
		Pos:      pos,
		Filename: l.currentFile,
	}
}

func (l *Lexer) readNumber() token.Token {
	var number bytes.Buffer
	pos := l.curPosition()
	base := ""
	tokenType := token.INT

	if l.currentCh == '0' && l.peekCh == 'x' {
		base = "0x"
		pos.Column-- // Correct for initial 0x
		l.readRune()
		l.readRune()
	}

	if l.currentCh == 'x' {
		base = "0x"
		pos.Column--
		l.readRune()
	}

	if l.currentCh == '.' {
		l.readRune()
		return token.Token{
			Literal:  "Invalid float literal",
			Type:     token.ILLEGAL,
			Pos:      pos,
			Filename: l.currentFile,
		}
	}

	for isDigit(l.currentCh) || isHexDigit(l.currentCh) {
		if l.currentCh == '.' {
			if tokenType != token.INT {
				return token.Token{
					Literal:  "Invalid float literal",
					Type:     token.ILLEGAL,
					Pos:      pos,
					Filename: l.currentFile,
				}
			}

			tokenType = token.FLOAT
		}

		number.WriteRune(l.currentCh)
		l.readRune()
	}

	return token.Token{
		Literal:  base + number.String(),
		Type:     token.TokenType(tokenType),
		Pos:      pos,
		Filename: l.currentFile,
	}
}

// TODO: maybe discard instead of reading?
func (l *Lexer) readSingleLineComment() token.Token {
	var comment bytes.Buffer
	pos := l.curPosition()

	if l.currentCh == '/' {
		pos.Column--
	}
	l.readRune()

	for l.currentCh != '\n' && l.currentCh != '0' {
		comment.WriteRune(l.currentCh)
		l.readRune()
	}

	return token.Token{
		Literal:  strings.TrimSpace(comment.String()),
		Type:     token.COMMENT,
		Pos:      pos,
		Filename: l.currentFile,
	}
}

func (l *Lexer) readMultiLineComment() token.Token {
	var comment bytes.Buffer
	pos := l.curPosition()
	pos.Column--
	l.readRune()

	for l.currentCh != 0 {
		if l.currentCh == '*' && l.peekChar() == '/' {
			l.readRune()
			break
		}

		if l.currentCh == '\n' {
			l.resetPos()
		}

		comment.WriteRune(l.currentCh)
		l.readRune()
	}

	return token.Token{
		Literal:  comment.String(),
		Type:     token.COMMENT,
		Pos:      pos,
		Filename: l.currentFile,
	}
}

func (l *Lexer) trimWhiteSpace() {
	for l.isWhiteSpace(l.currentCh) {
		if l.currentCh == '\n' {
			l.resetPos()
		}
		l.readRune()
	}
}

func (l *Lexer) trimWhiteSpaceNotNewLine() {
	for l.currentCh != '\n' && l.isWhiteSpace(l.currentCh) {
		l.readRune()
	}
}

func (l *Lexer) newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{
		Literal:  string(ch),
		Type:     tokenType,
		Pos:      l.curPosition(),
		Filename: l.currentFile,
	}
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || unicode.IsLetter(ch)
}

func isIdent(ch rune) bool {
	return isLetter(ch) || (ch != '.' && isDigit(ch))
}

func isDigit(ch rune) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func isHexDigit(ch rune) bool {
	return ('a' <= ch && xh <= 'f') || ('A' <= ch && xh <= 'F')
}

func (l *Lexer) isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) lastTokenWas(types ...token.TokenType) bool {
	for _, t := range types {
		if l.lastToken.Type == t {
			return true
		}
	}
	return false
}
