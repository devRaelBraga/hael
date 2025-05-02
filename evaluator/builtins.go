package evaluator

import (
	"hael/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, text := range args {
				println(text.Inspect())
			}

			return NULL
		},
	},

	// "map": &object.Builtin{
	// 	Fn: func(args ...object.Object) object.Object {
	// 		// if len(args) != 2 {
	// 		// 	return newError("wrong number of arguments. got=%d, want=2", len(args))
	// 		// }

	// 		if args[0].Type() != object.ARRAY_OBJ {
	// 			return newError("argument to `map` must be ARRAY, got %s", args[0].Type())
	// 		}
	// 		// if args[1].Type() != object.FUNCTION_OBJ {
	// 		// 	return newError("2nd argument to `map` must be FUNCTION, got %s", args[1].Type())
	// 		// }

	// 		arr := args[0].(*object.Array)
	// 		// fn := args[1].(*object.Function)

	// 		// for i, element := range arr.Elements {
	// 		// 	// arr.Elements[i] = &object.Integer{Value: element.(*object.Integer).Value + 1}
	// 		// 	fmt.Println(fn)
	// 		// }

	// 		length := len(arr.Elements)

	// 		newElements := make([]object.Object, length)
	// 		copy(newElements, arr.Elements)

	// 		return &object.Array{Elements: newElements}
	// 	},
	// },
}
