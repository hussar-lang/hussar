package ast_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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

func TestInterfaceImplementation(t *testing.T) {
	cases := []struct {
		node              interface{}
		expectedInterface interface{}
	}{
		{
			node:              new(ast.Program),
			expectedInterface: (*ast.Node)(nil),
		},
		{
			node:              new(ast.VarStatement),
			expectedInterface: (*ast.Statement)(nil),
		},
		{
			node:              new(ast.ReturnStatement),
			expectedInterface: (*ast.Statement)(nil),
		},
		{
			node:              new(ast.ExpressionStatement),
			expectedInterface: (*ast.Statement)(nil),
		},
		{
			node:              new(ast.Identifier),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.IntegerLiteral),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.StringLiteral),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.PrefixExpression),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.InfixExpression),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.Boolean),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.IfExpression),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.WhileExpression),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.BlockStatement),
			expectedInterface: (*ast.Statement)(nil),
		},
		{
			node:              new(ast.FunctionLiteral),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.CallExpression),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.ArrayLiteral),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.IndexExpression),
			expectedInterface: (*ast.Expression)(nil),
		},
		{
			node:              new(ast.ExitLiteral),
			expectedInterface: (*ast.Expression)(nil),
		},
	}

	for _, tc := range cases {
		assert.Implements(t, tc.expectedInterface, tc.node, "it should implement the interface")
	}
}
