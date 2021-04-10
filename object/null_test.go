package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.io/lang/object"
)

func TestNullType(t *testing.T) {
	val := &object.Null{}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("NULL"), valType, "the type should match")
}

func TestNullInspect(t *testing.T) {
	val := &object.Null{}
	valInspect := val.Inspect()

	assert.Equal(t, "null", valInspect, "the string representation should match")
}
