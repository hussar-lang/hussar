package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.io/lang/ast"
	"hussar.io/lang/object"
)

func TestFunctionType(t *testing.T) {
	val := &object.Function{
		Parameters: nil,
		Body:       nil,
		Env:        nil,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("FUNCTION"), valType, "the type should match")
}

func TestFunctionInspect(t *testing.T) {
	// TODO: find a way to test this better by feeding in a valid function but also not coupling it too much to ast. Maybe with mocking?
	val := &object.Function{
		Parameters: []*ast.Identifier{},
		Body:       &ast.BlockStatement{},
		Env:        &object.Environment{},
	}
	valInspect := val.Inspect()

	assert.Equal(t, "fn() {\n\n}", valInspect, "the string representation should match")
}
