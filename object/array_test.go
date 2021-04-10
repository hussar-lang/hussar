package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestArrayType(t *testing.T) {
	val := &object.Array{
		Elements: nil,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("ARRAY"), valType, "the type should match")
}

func TestArrayInspect(t *testing.T) {
	val := &object.Array{
		Elements: []object.Object{
			&object.String{
				Value: "foo bar",
			},
			&object.String{
				Value: "foo fighters",
			},
		},
	}
	valInspect := val.Inspect()

	assert.Equal(t, "[foo bar, foo fighters]", valInspect, "the string representation should match")
}
