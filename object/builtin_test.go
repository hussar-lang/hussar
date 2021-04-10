package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestBuiltinType(t *testing.T) {
	val := &object.Builtin{
		Fn: nil,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("BUILTIN"), valType, "the type should match")
}

func TestBuiltinInspect(t *testing.T) {
	val := &object.Builtin{
		Fn: nil,
	}
	valInspect := val.Inspect()

	assert.Equal(t, "builtin function", valInspect, "the string representation should match")
}
