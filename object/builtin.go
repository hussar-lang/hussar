package object

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the type of object represented
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Inspect returns a string representation of the value
func (b *Builtin) Inspect() string { return "builtin function" }
