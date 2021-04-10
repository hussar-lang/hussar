package ast_test

import (
	"testing"

	"hussar.dev/lang/ast"
	"hussar.dev/lang/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.VarStatement{
				Token: token.Token{Type: token.Var, Literal: "var"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.Identifier, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &ast.Identifier{
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
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Token: token.Token{Type: token.Return, Literal: "return"},
				ReturnValue: &ast.PrefixExpression{
					Token:    token.Token{Type: token.Bang, Literal: "!"},
					Operator: "!",
					Right: &ast.Identifier{
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
