package object

import "fmt"

type Boolean struct {
	Value bool
}

// Type returns the type of object represented
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect returns a string representation of the value
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
