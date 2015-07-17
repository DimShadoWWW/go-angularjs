package angularjs

import "fmt"
import "reflect"
import "github.com/gopherjs/gopherjs/js"

func typeOf(cb interface{}) (reflect.Type, bool) {
	typ := reflect.TypeOf(cb)

	if typ.Kind() == reflect.Ptr {
		if meth, ok := typ.MethodByName("New"); ok {
			return meth.Type, true
		}
	}

	return typ, false
}

func makeConstructor(cb interface{}) func(this *js.Object, args []*js.Object) interface{} {
	typ := reflect.TypeOf(cb)

	newMethod, ok := typ.MethodByName("New")
	if !ok && typ.Kind() != reflect.Func {
		panic(fmt.Sprintf("Type %q doesn't have New method.", typ))
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	return func(jsThis *js.Object, jsArgs []*js.Object) interface{} {
		this := reflect.New(typ).Elem()

		objField := this.FieldByName("Object")
		objField.Set(reflect.ValueOf(jsThis))

		var args []reflect.Value
		args = append(args, this.Addr())

		for i, jsArg := range jsArgs {
			argTyp := newMethod.Type.In(i + 1)
			if argTyp.Kind() == reflect.Ptr {
				argTyp = argTyp.Elem()
			}

			arg := reflect.New(argTyp)
			arg.Elem().FieldByName("Object").Set(reflect.ValueOf(jsArg))

			if meth, ok := arg.Type().MethodByName("Init"); ok {
				//println("Calling Init for arg ", arg.Elem().Type().String())

				meth.Func.Call([]reflect.Value{arg})
			}

			if newMethod.Type.In(i+1).Kind() == reflect.Ptr {
				args = append(args, arg)
			} else {
				args = append(args, arg.Elem())
			}
		}

		newMethod.Func.Call(args)

		return nil
	}
}

func funcReturn(cb interface{}, ret []interface{}) error {
	typ, _ := typeOf(cb)

	if typ.NumOut() != len(ret) {
		return fmt.Errorf("Wrong function signature -- Expected %q, got %q", ret, typ)
	}

	for i := 0; i < typ.NumOut(); i++ {
		if typ.Out(i).String() != reflect.TypeOf(ret[i]).String() {
			return fmt.Errorf("Wrong function signature -- Expected %q, got %q", ret, typ)
		}
	}

	return nil
}

func buildParams(cb interface{}, name string) (js.S, error) {
	arg := js.S{}

	typ, cls := typeOf(cb)
	if typ.Kind() != reflect.Func {
		return nil, fmt.Errorf("Expected a function, got %q", typ)
	}

	i := 0

	if cls {
		i = 1
	}

	for ; i < typ.NumIn(); i++ {
		argTyp := typ.In(i)
		if argTyp.Kind() != reflect.Struct {
			argTyp = argTyp.Elem()
		}

		param, ok := argTyp.FieldByName("Object")

		if !ok {
			return nil, fmt.Errorf("Invalid paramater type %q (Missing Object)", argTyp)
		}

		if param.Tag.Get(name) == "" {
			if param.Tag.Get("js") == "" {
				return nil, fmt.Errorf("Invalid paramater type %q (Missing tag %s)", argTyp, name)
			} else {
				arg = append(arg, param.Tag.Get("js"))
			}
		} else {
			arg = append(arg, param.Tag.Get(name))
		}
	}

	if cls {
		arg = append(arg, js.MakeFunc(makeConstructor(cb)))
	} else {
		arg = append(arg, cb)
	}

	return arg, nil
}
