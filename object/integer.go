package object

import "fmt"

type Integer struct {
	Value int64
}

// Type returns the type of object represented
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Inspect returns a string representation of the value
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
