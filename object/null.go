package object

type Null struct{}

// Type returns the type of object represented
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Inspect returns a string representation of the value
func (n *Null) Inspect() string { return "null" }
