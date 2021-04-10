package object_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"hussar.dev/lang/object"
)

func TestErrorInterface(t *testing.T) {
	assert.Implements(t, (*object.Object)(nil), new(object.Error), "it should implement the interface")
}

func TestErrorType(t *testing.T) {
	val := &object.Error{
		Severity: "",
		Message:  "",
	}
	valType := val.Type()

	assert.Equal(t, object.ObjectType("ERROR"), valType, "the type should match")
}

func TestErrorInspect(t *testing.T) {
	cases := []struct {
		errorObj         object.Error
		expectedInspect  string
		expectedSeverity string
		expectedMessage  string
	}{
		{
			errorObj: object.Error{
				Severity: "warn",
				Message:  "test message",
			},
			expectedInspect:  "\x1b[40m\x1b[33m\x1b[1m[WARN]  \x1b[22m\x1b[49m\x1b[39mtest message",
			expectedSeverity: "warn",
			expectedMessage:  "test message",
		},
		{
			errorObj: object.Error{
				Severity: "error",
				Message:  "another test message",
			},
			expectedInspect:  "\x1b[40m\x1b[31m\x1b[1m[ERROR] \x1b[22m\x1b[49m\x1b[39manother test message",
			expectedSeverity: "error",
			expectedMessage:  "another test message",
		},
		{
			errorObj: object.Error{
				Severity: "otherValueDefaultPrint",
				Message:  "yet another test message",
			},
			expectedInspect:  "\x1b[40m\x1b[31m\x1b[1m[ERROR] \x1b[22m\x1b[49m\x1b[39myet another test message",
			expectedSeverity: "otherValueDefaultPrint",
			expectedMessage:  "yet another test message",
		},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.expectedInspect, tc.errorObj.Inspect(), "the string representation should match")
		assert.Equal(t, tc.expectedSeverity, tc.errorObj.Severity, "the severity should match")
		assert.Equal(t, tc.expectedMessage, tc.errorObj.Message, "the error message should match")
	}
}
