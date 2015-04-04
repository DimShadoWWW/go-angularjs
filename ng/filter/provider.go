package filter

import "github.com/gopherjs/gopherjs/js"

type Provider struct {
	*js.Object `ajs-provider:"$filterProvider"`
}

// Registers a filter function with $filterProvider
func (p *Provider) Register(name string) {
	p.Call("register", name, js.S{name, func(v interface{}) interface{} {
		return v
	}})
}
