package angularjs

import "github.com/gopherjs/gopherjs/js"

type JQueryElement struct{
	*js.Object `ajs-service:"element"`
}

func ElementById(id string) *JQueryElement {
	return &JQueryElement{js.Global.Get("angular").Call("element", js.Global.Get("document").Call("getElementById", id))}
}

func (e *JQueryElement) Text() string {
	return e.Call("text").String()
}

func (e *JQueryElement) SetText(val string) {
	e.Call("text", val)
}

func (e *JQueryElement) Prop(name string) *js.Object {
	return e.Call("prop", name)
}

func (e *JQueryElement) SetProp(name, value interface{}) {
	e.Call("prop", name, value)
}

func (e *JQueryElement) On(events string, handler func(*Event)) {
	e.Call("on", events, func(e *js.Object) {
		handler(&Event{Object: e})
	})
}

func (e *JQueryElement) Attr(name string) *js.Object {
	return e.Call("attr", name)
}

func (e *JQueryElement) SetAttr(name string, value interface{}) {
	e.Call("attr", name, value)
}

func (e *JQueryElement) Val() *js.Object {
	return e.Call("val")
}

func (e *JQueryElement) SetVal(value interface{}) {
	e.Call("val", value)
}

func (e *JQueryElement) Watch(prop string, cb interface{}) {
	e.Call("scope").Call("$watch", prop, cb)
}
