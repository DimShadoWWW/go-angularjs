package route

import "github.com/gopherjs/gopherjs/js"

type Provider struct {
	*js.Object `ajs-provider:"$routeProvider"`
}

func buildRoute(defs []RouteDef) *route {
	r := &route{
		Object: js.Global.Get("Object").New(),
	}

	for _, def := range defs {
		def(r)
	}

	return r
}

func (p *Provider) When(path string, defs ...RouteDef) {
	p.Call("when", path, buildRoute(defs))
}

func (p *Provider) Otherwise(defs ...RouteDef) {
	p.Call("otherwise", buildRoute(defs))
}
