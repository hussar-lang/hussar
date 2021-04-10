package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestBooleanInterface(t *testing.T) {
	assert.Implements(t, (*object.Object)(nil), new(object.Boolean), "it should implement the interface")
}

func TestBooleanType(t *testing.T) {
	val := &object.Boolean{
		Value: true,
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("BOOLEAN"), valType, "the type should match")
}

func TestBooleanInspect(t *testing.T) {
	// Is this overkill? Yes
	// Does it work? Also yes
	cases := []struct {
		object   bool
		expected string
	}{
		{
			object:   false,
			expected: "false",
		},
		{
			object:   true,
			expected: "true",
		},
	}

	for _, tc := range cases {

		val := &object.Boolean{
			Value: tc.object,
		}
		valInspect := val.Inspect()

		assert.Equal(t, tc.expected, valInspect, "the string representation should match")

	}
}
