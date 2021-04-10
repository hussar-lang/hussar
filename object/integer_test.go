package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestIntegerType(t *testing.T) {
	val := &object.Integer{
		Value: 512,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("INTEGER"), valType, "the type should match")
}

func TestIntegerInspect(t *testing.T) {
	val := &object.Integer{
		Value: 922,
	}
	valInspect := val.Inspect()

	assert.Equal(t, "922", valInspect, "the string representation should match")
}
