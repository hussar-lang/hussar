package ast

import (
	"testing"

	"hussar.io/lang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
				Token: token.Token{Type: token.Var, Literal: "var"},
				Name: &Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "var myVar = anotherVar;" {
		t.Errorf("program.String() wrong. Got: %q", program.String())
	}
}

// Thought this might be useful
func TestTokenLiteralString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&ReturnStatement{
				Token: token.Token{Type: token.Return, Literal: "return"},
				ReturnValue: &PrefixExpression{
					Token:    token.Token{Type: token.Bang, Literal: "!"},
					Operator: "!",
					Right: &Identifier{
						Token: token.Token{Type: token.Identifier, Literal: "myVar"},
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
