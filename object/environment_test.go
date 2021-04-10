package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestEnvironmentCreation(t *testing.T) {
	env := object.NewEnvironment()

	// There's only so much we can do when we can't access the unexported values
	assert.NotNil(t, env, "it should exist")
}

func TestEnclosedEnvironmentCreation(t *testing.T) {
	outer := object.NewEnvironment()
	inner := object.NewEnclosedEnvironment(outer)

	// There's only so much we can do when we can't access the unexported values
	assert.NotNil(t, outer, "it should exist")
	assert.NotNil(t, inner, "it should exist")
}

func TestEnvironmentRetrieval(t *testing.T) {
	cases := []struct {
		name          string
		object        object.Object
		expectedValue interface{}
	}{
		{"testString", &object.String{
			Value: "foo",
		}, "foo"},
		{"testBoolean", &object.Boolean{
			Value: false,
		}, "false"},
		{"testInt", &object.Integer{
			Value: 552,
		}, "552"},
	}

	env := object.NewEnvironment()

	for _, tc := range cases {
		env.Set(tc.name, tc.object)
	}

	for _, tc := range cases {
		val, ok := env.Get(tc.name)

		assert.True(t, ok, "it should be able to retrieve the value")
		assert.Equal(t, tc.expectedValue, val.Inspect(), "the retrieved value should match the entered value")
	}
}

func TestEnclosedEnvironmentRetrieval(t *testing.T) {
	outer := object.NewEnvironment()
	middle := object.NewEnclosedEnvironment(outer)
	inner := object.NewEnclosedEnvironment(middle)

	outer.Set("outerString", &object.String{
		Value: "foo",
	})
	middle.Set("middleString", &object.String{
		Value: "bar",
	})
	inner.Set("innerString", &object.String{
		Value: "baz",
	})

	val, ok := inner.Get("outerString")
	assert.True(t, ok, "inner should be able to retrieve the value from outer")
	assert.Equal(t, "foo", val.Inspect(), "the retrieved value should match the entered value")

	val, ok = inner.Get("middleString")
	assert.True(t, ok, "inner should be able to retrieve the value from middle")
	assert.Equal(t, "bar", val.Inspect(), "the retrieved value should match the entered value")

	val, ok = inner.Get("innerString")
	assert.True(t, ok, "inner should be able to retrieve the value from inner")
	assert.Equal(t, "baz", val.Inspect(), "the retrieved value should match the entered value")

	val, ok = middle.Get("outerString")
	assert.True(t, ok, "middle should be able to retrieve the value from outer")
	assert.Equal(t, "foo", val.Inspect(), "the retrieved value should match the entered value")

	val, ok = middle.Get("middleString")
	assert.True(t, ok, "middle should be able to retrieve the value from middle")
	assert.Equal(t, "bar", val.Inspect(), "the retrieved value should match the entered value")

	val, ok = middle.Get("innerString")
	assert.False(t, ok, "middle should not be able to retrieve the value from inner")
}

func TestEnclosedEnvironmentShadowing(t *testing.T) {
	outer := object.NewEnvironment()
	inner := object.NewEnclosedEnvironment(outer)

	outer.Set("testBoolean", &object.Boolean{
		Value: true,
	})

	inner.Set("testBoolean", &object.Boolean{
		Value: false,
	})

	val, ok := inner.Get("testBoolean")

	assert.True(t, ok, "it should be able to retrieve the value from the inner environment")
	assert.Equal(t, "false", val.Inspect(), "the enclosed environment should override the parent")

	val, ok = outer.Get("testBoolean")

	assert.True(t, ok, "it should be able to retrieve the value from the outer environment")
	assert.Equal(t, "true", val.Inspect(), "the outer environment should stay intact")
}
