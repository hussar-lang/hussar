package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
}

// Type returns the type of object represented
func (a *Array) Type() ObjectType { return ARRAY_OBJ }

// Inspect returns a string representation of the value
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
