package token

// Improvement: Use byte for better performance (pg12)
type TokenType string

type Position struct {
	Line, Col int
}

// Improvement: include filename, line and character place to ease debugging
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
	ILLEGAL: "ILLEGAL"
	EOF     = "EOF"
	COMMENT     = "COMMENT"
	EOL: "EOL"

	// Identifiers + literals
	IDENT  = "IDENT"
	INT    = "INT"
	FLOAT: "FLOAT"
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// Delimitors
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	WHILE    = "WHILE"
	RETURN   = "RETURN"
	EXIT     = "EXIT"
	NIL: "NIL"
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"while":  WHILE,
	"return": RETURN,
	"exit":   EXIT,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
