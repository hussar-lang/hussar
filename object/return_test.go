package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestReturnInterface(t *testing.T) {
	assert.Implements(t, (*object.Object)(nil), new(object.ReturnValue), "it should implement the interface")
}

func TestReturnType(t *testing.T) {
	val := &object.ReturnValue{
		Value: nil,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("RETURN_VALUE"), valType, "the type should match")
}

func TestReturnInspect(t *testing.T) {
	val := &object.ReturnValue{
		Value: &object.Integer{
			Value: 5,
		},
	}
	valInspect := val.Inspect()

	assert.Equal(t, "5", valInspect, "the string representation should match")
}
