package token

// Improvement: Use byte for better performance (pg12)
//go:generate stringer -type=TokenType
type TokenType int

type Position struct {
	Line, Col int
}

// Improvement: include filename, line and character place to ease debugging (this is implemented in some form now - modify to make use of it).
type Token struct {
	Type     TokenType
	Literal  string
	Pos      Position
	Filename string
}

const (
	Illegal TokenType = iota + 1
	EOF
	Comment
	EOL

	// Identifiers + literals
	Identifier
	Integer
	Float
	String

	// Operators
	Assign
	Plus
	Minus
	Bang
	Asterisk
	Slash

	LessThan
	GreaterThan

	Equal
	NotEqual

	// Delimitors
	Comma
	SemiColon

	LParen
	RParen
	LBrace
	RBrace
	LBracket
	RBracket

	// Keywords
	Function
	Var
	True
	False
	Nil
	If
	Else
	While
	Return
	Exit
)

// TODO: use the above and apply into this + change all other occurences
var tokens = [...]string{
	Illegal: "ILLEGAL",
	EOF:     "EOF",
	Comment: "COMMENT",
	EOL:     "EOL",

	// Identifiers + literals
	Identifier: "IDENT",
	Integer:    "INT",
	Float:      "FLOAT",
	String:     "STRING",

	// Operators
	Assign:   "=",
	Plus:     "+",
	Minus:    "-",
	Bang:     "!",
	Asterisk: "*",
	Slash:    "/",

	LessThan:    "<",
	GreaterThan: ">",

	Equal:    "==",
	NotEqual: "!=",

	// Delimitors
	Comma:     ",",
	SemiColon: ";",

	LParen:   "(",
	RParen:   ")",
	LBrace:   "{",
	RBrace:   "}",
	LBracket: "[",
	RBracket: "]",

	// Keywords
	Function: "FUNCTION",
	Var:      "VAR",
	True:     "TRUE",
	False:    "FALSE",
	If:       "IF",
	Else:     "ELSE",
	While:    "WHILE",
	Return:   "RETURN",
	Exit:     "EXIT",
	Nil:      "NIL",
}

var keywords = map[string]TokenType{
	"fn":     Function,
	"var":    Var,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"while":  While,
	"return": Return,
	"exit":   Exit,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Identifier
}
