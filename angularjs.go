package angularjs

import "github.com/gopherjs/gopherjs/js"

type Module struct{ *js.Object }

func (m *Module) NewController(name string, constructor func(scope *Scope)) {
	m.Call("controller", name, js.S{"$scope", func(scope *js.Object) {
		constructor(&Scope{Object: scope})
	}})
}

func (m *Module) Service(name string, cls interface{}) error {
	var err error

	// if err := funcReturn(cls, nil); err != nil {
	// 	return err
	// }

	arg, err := buildParams(cls, "ajs-service")
	if err != nil {
		println("Error building controller params: ", err.Error(), "controller: ", name)

		return err
	}

	m.Call("service", name, arg)

	return nil
}

func (m *Module) Controller(name string, cls interface{}) error {
	var err error

	// if err := funcReturn(cls, nil); err != nil {
	// 	return err
	// }

	arg, err := buildParams(cls, "ajs-service")
	if err != nil {
		println("Error building controller params: ", err.Error(), "controller: ", name)

		return err
	}

	m.Call("controller", name, arg)
	js.Global.Set(name, js.MakeFunc(makeConstructor(cls)))

	return nil
}

func (m *Module) Config(cb interface{}) error {
	if err := funcReturn(cb, nil); err != nil {
		return err
	}

	arg, err := buildParams(cb, "ajs-provider")
	if err != nil {
		println("Error building config params:", err.Error())

		return err
	}

	m.Call("config", arg)

	return nil
}

func (m *Module) Run(cb interface{}) error {
	if err := funcReturn(cb, nil); err != nil {
		return err
	}

	arg, err := buildParams(cb, "ajs-service")
	if err != nil {
		println("Error building run params:", err.Error())

		return err
	}

	m.Call("run", arg)

	return nil
}

func (m *Module) Value(name string, val interface{}) error {
	m.Call("value", name, val)

	return nil
}

func (m *Module) Filter(name string, val interface{}) error {
	if err := funcReturn(val, []interface{}{""}); err != nil {
		return err
	}

	if err := m.Value(name, val); err != nil {
		return err
	}

	m.Call("filter", name, js.S{name, func(v interface{}) interface{} {
		return v
	}})

	return nil
}

func (m *Module) Directive(name string, cb interface{}) error {
	arg, err := buildParams(cb, "ajs-service")
	if err != nil {
		return err
	}

	m.Call("directive", name, arg)

	return nil
}

type Scope struct {
	*js.Object `ajs-service:"$scope"`
}

func (s *Scope) Apply(f func()) {
	s.Call("$apply", f)
}

func (s *Scope) EvalAsync(f func()) {
	s.Call("$evalAsync", f)
}

type Event struct {
	*js.Object

	KeyCode int `js:"keyCode"`
}

func (e *Event) PreventDefault() {
	e.Call("preventDefault")
}

var angularModule *js.Object

func NewModule(name string, requires []string) *Module {
	// defer println("Failed to require:", recover())

	// require := js.Global.Get("require")
	// if require != js.Undefined {
	// 	println("Requiring angular...")

	// 	angularModule = require.Invoke("angular")
	// }

	return &Module{js.Global.Get("angular").Call("module", name, requires)}
}
