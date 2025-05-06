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
	"pop": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `pop` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) == 0 {
				return newError("cannot pop from empty array")
			}

			lastElement := arr.Elements[len(arr.Elements)-1]

			newElements := make([]object.Object, len(arr.Elements)-1)
			copy(newElements, arr.Elements[:len(arr.Elements)-1])
			arr.Elements = newElements

			return lastElement
		},
	},
	"type": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("type expects one argument")
			}
			return &object.String{Value: string(args[0].Type())}
		},
	},
	"slice": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}

			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `slice` must be ARRAY, got %s", args[0].Type())
			}

			if args[1].Type() != object.INTEGER_OBJ {
				return newError("second argument to `slice` must be INTEGER, got %s", args[1].Type())
			}

			end := int64(len(args[0].(*object.Array).Elements))

			if len(args) == 3 {
				if args[2].Type() != object.INTEGER_OBJ {
					return newError("third argument to `slice` must be INTEGER, got %s", args[2].Type())
				}
				end = args[2].(*object.Integer).Value
			}

			arr := args[0].(*object.Array)
			start := args[1].(*object.Integer).Value

			if start < 0 || end > int64(len(arr.Elements)) || start > end {
				return newError("invalid slice indices: start=%d, end=%d", start, end)
			}

			newElements := make([]object.Object, end-start)
			copy(newElements, arr.Elements[start:end])

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
