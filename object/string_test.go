package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestStringType(t *testing.T) {
	val := &object.String{
		Value: "foo",
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("STRING"), valType, "the type should match")
}

func TestStringInspect(t *testing.T) {
	val := &object.String{
		Value: "foo",
	}
	valInspect := val.Inspect()

	assert.Equal(t, "foo", valInspect, "the string representation should match")
}
