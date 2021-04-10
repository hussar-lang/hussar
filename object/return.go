package object

type ReturnValue struct {
	Value Object
}

// Type returns the type of object represented
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect returns a string representation of the value
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }
