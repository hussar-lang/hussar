package object

import (
	"bytes"

	"hussar.io/lang/ast"
)

type Exit struct {
	ExitCode *ast.Expression
}

func (e *Exit) Type() ObjectType { return EXIT_OBJ }
func (e *Exit) Inspect() string {
	var out bytes.Buffer

	out.WriteString("exit()")

	return out.String()
}
