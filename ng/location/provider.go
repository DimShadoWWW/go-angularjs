package location

import "github.com/gopherjs/gopherjs/js"

type Provider struct {
	*js.Object `ajs-provider:"$locationProvider"`
}

func (p *Provider) SetHtml5Mode(val bool) {
	p.Call("html5Mode", val)
}

func (p *Provider) GetHtml5Mode() bool {
	return p.Call("html5Mode").Bool()
}
