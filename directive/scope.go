package directive

import "github.com/gopherjs/gopherjs/js"

type ScopeValue string

func Scope(name string, value ScopeValue) directiveMod {
	return func(d *directive) {
		if d.Get("scope") == js.Undefined {
			d.Set("scope", js.Global.Get("Object").New())
		}

		d.Get("scope").Set(name, value)

		println("Scope:", d.Get("Scope"))
	}
}

func ScopeAttr(attr string) ScopeValue {
	return ScopeValue("@" + attr)
}
