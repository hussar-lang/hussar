package object

type String struct {
	Value string
}

// Type returns the type of object represented
func (s *String) Type() ObjectType { return STRING_OBJ }

// Inspect returns a string representation of the value
func (s *String) Inspect() string { return s.Value }
