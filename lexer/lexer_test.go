package lexer

import (
	"strings"
	"testing"

	"hussar.io/lang/token"
)

func TestNextToken(t *testing.T) {
	input := `var five = 5;
	var ten = 10;

	var add = fn(x, y) {
		x + y;$
	};

	var result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	//comment

	10 == 10;
	10 != 9; // Inline comment
	"foobar"
	"foo bar"
	[1, 2];
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.Var, "var"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.SemiColon, ";"},
		{token.Var, "var"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.SemiColon, ";"},
		{token.Var, "var"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.SemiColon, ";"},
		{token.Illegal, "$"},
		{token.RBrace, "}"},
		{token.SemiColon, ";"},
		{token.Var, "var"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.LParen, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RParen, ")"},
		{token.SemiColon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Integer, "5"},
		{token.SemiColon, ";"},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.GreaterThan, ">"},
		{token.Integer, "5"},
		{token.SemiColon, ";"},
		{token.If, "if"},
		{token.LParen, "("},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.SemiColon, ";"},
		{token.RBrace, "}"},
		{token.Else, "else"},
		{token.LBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.SemiColon, ";"},
		{token.RBrace, "}"},
		{token.Integer, "10"},
		{token.Equal, "=="},
		{token.Integer, "10"},
		{token.SemiColon, ";"},
		{token.Integer, "10"},
		{token.NotEqual, "!="},
		{token.Integer, "9"},
		{token.SemiColon, ";"},
		{token.String, "foobar"},
		{token.String, "foo bar"},
		{token.LBracket, "["},
		{token.Integer, "1"},
		{token.Comma, ","},
		{token.Integer, "2"},
		{token.RBracket, "]"},
		{token.SemiColon, ";"},
		{token.EOF, ""},
	}

	l := New(strings.NewReader(input))

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - literal wrong. expected: %q (%s), got: %q (%s)", i, tt.expectedLiteral, string(tt.expectedType), tok.Literal, string(tok.Type))
		}
	}
}
