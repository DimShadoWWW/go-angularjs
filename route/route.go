package route

import "github.com/gopherjs/gopherjs/js"

type route struct {
	*js.Object

	Controller  string `js:"controller"`
	Template    string `js:"template"`
	TemplateUrl string `js:"templateUrl"`
	RedirectTo  string `js:"redirectTo"`
}

type RouteDef func(*route)

func (p *Provider) Controller(ctrl string) func(*route) {
	return func(r *route) {
		r.Controller = ctrl
	}
}

func (p *Provider) Redirect(tgt string) func(*route) {
	return func(r *route) {
		r.RedirectTo = tgt
	}
}

func (p *Provider) TemplateUrl(tgt string) func(*route) {
	return func(r *route) {
		r.TemplateUrl = tgt
	}
}
