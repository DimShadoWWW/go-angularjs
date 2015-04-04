package angularjs

import "fmt"
import "reflect"
import "github.com/gopherjs/gopherjs/js"

func funcReturn(cb interface{}, ret []interface{}) error {
	typ := reflect.TypeOf(cb)

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

	typ := reflect.TypeOf(cb)
	if typ.Kind() != reflect.Func {
		return nil, fmt.Errorf("Expected a function, got %q", typ)
	}

	for i := 0; i < typ.NumIn(); i++ {
		if typ.In(i).Kind() != reflect.Struct {
			return nil, fmt.Errorf("Invalid paramater type %q (Not Struct)", typ.In(i))
		}

		param, ok := typ.In(i).FieldByName("Object")

		if !ok {
			return nil, fmt.Errorf("Invalid paramater type %q (Missing Object)", typ.In(i))
		}

		if param.Tag.Get(name) == "" {
			return nil, fmt.Errorf("Invalid paramater type %q (Missing tag %s)", typ.In(i), name)
		}

		arg = append(arg, param.Tag.Get(name))
	}

	arg = append(arg, cb)

	return arg, nil
}
