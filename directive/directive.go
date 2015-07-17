package directive

import "github.com/gopherjs/gopherjs/js"
import "github.com/AmandaCameron/go-angularjs"

type DirectiveLink func(angularjs.Scope, angularjs.JQueryElement)

type directive struct {
	*js.Object

	Restrict    DirectivePos `js:"restrict"`
	TemplateUrl string       `js:"templateUrl"`
	Template    string       `js:"template"`
	Transclude  bool         `js:"transclude"`
	Controller  interface{}  `js:"controller"`
}

type directiveMod func(*directive)

func New(mods ...directiveMod) *js.Object {
	d := &directive{
		Object: js.Global.Get("Object").New(),
	}

	for _, mod := range mods {
		mod(d)
	}

	return d.Object
}
