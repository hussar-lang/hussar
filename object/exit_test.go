package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestExitType(t *testing.T) {
	val := &object.Exit{
		ExitCode: nil,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("EXIT"), valType, "the type should match")
}

func TestExitInspect(t *testing.T) {
	// TODO: find a way to test this better by feeding in a valid ast.expression but also not coupling it too much to ast. Maybe with mocking?
	val := &object.Exit{
		ExitCode: nil,
	}
	valInspect := val.Inspect()

	assert.Equal(t, "exit()", valInspect, "the string representation should match")
}
