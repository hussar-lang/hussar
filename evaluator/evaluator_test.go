package evaluator

import (
	"github.com/kscarlett/kmonkey/lexer"
	"github.com/kscarlett/kmonkey/object"
	"github.com/kscarlett/kmonkey/parser"

	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("Object is of incorrect type. Expected: %s, got: %T (%+v)", "object.Integer", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("Object has an unexpected value. Expected: %d, got: %d", expected, result.Value)
		return false
	}

	return true
}
