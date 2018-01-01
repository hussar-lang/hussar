package evaluator

import (
	"github.com/kscarlett/kmonkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("error", "Wrong number of arguments passed to `len`. Expected: %d, got: %d", 1, len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("error", "Argument %s is not supported by `len`", args[0].Type())
			}
		},
	},
}
