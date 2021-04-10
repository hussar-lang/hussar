package object

import (
	"bytes"

	"hussar.io/lang/ast"
)

type Exit struct {
	ExitCode *ast.Expression
}

// Type returns the type of object represented
func (e *Exit) Type() ObjectType { return EXIT_OBJ }

// Inspect returns a string representation of the value
func (e *Exit) Inspect() string {
	var out bytes.Buffer

	out.WriteString("exit()")

	return out.String()
}
