package object

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new environment wrapped in the outer environment
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

// Environment holds a scope for variables and objects
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns the requested object from the current or outer environment
func (e *Environment) Get(name string) (object Object, ok bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets the given object in the environment for the given name
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
