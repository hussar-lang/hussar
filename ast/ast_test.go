package ast

import (
	"testing"

	"github.com/hussar-lang/hussar/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got: %q", program.String())
	}
}

// Thought this might be useful
func TestTokenLiteralString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &PrefixExpression{
					Token:    token.Token{Type: token.BANG, Literal: "!"},
					Operator: "!",
					Right: &Identifier{
						Token: token.Token{Type: token.IDENT, Literal: "myVar"},
						Value: "myVar",
					},
				},
			},
		},
	}

	if program.String() != "return (!myVar);" {
		t.Errorf("program.String() wrong. Got: %q", program.String())
	}
}
