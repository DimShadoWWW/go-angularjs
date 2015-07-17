package directive

import "github.com/AmandaCameron/go-angularjs"

func Link(cb DirectiveLink) directiveMod {
	return func(d *directive) {
		d.Set("link", cb)
	}
}

type DirectivePos string

const Element DirectivePos = "E"
const Attribute DirectivePos = "A"

func Allow(pos DirectivePos) directiveMod {
	return func(d *directive) {
		d.Restrict += pos
	}
}

func Controller(val interface{}) directiveMod {
	val, err := angularjs.Annotate(val, angularjs.AnnotServices)
	if err != nil {
		panic(err)
	}

	return func(d *directive) {
		d.Controller = val
	}
}

func TemplateUrl(url string) directiveMod {
	return func(d *directive) {
		d.TemplateUrl = url
	}
}

func Template(txt string) directiveMod {
	return func(d *directive) {
		d.Template = txt
	}
}

func Transclude() directiveMod {
	return func(d *directive) {
		d.Transclude = true
	}
}

func NoTransclude() directiveMod {
	return func(d *directive) {
		d.Transclude = false
	}
}
