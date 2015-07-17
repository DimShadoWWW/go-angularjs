package route

import "github.com/gopherjs/gopherjs/js"

type modFunc func(*routeConfig)

type routeConfig struct {
	*js.Object

	Path        string `js:"path"`
	Controller  string `js:"component"`
	RedirectTo  string `js:"redirectTo"`
	TemplateUrl string `js:"templateUrl"`
}

func Component(val string) modFunc {
	return modFunc(func(cfg *routeConfig) {
		cfg.Controller = val
	})
}

func Redirect(val string) modFunc {
	return modFunc(func(cfg *routeConfig) {
		cfg.RedirectTo = val
	})
}

func TemplateUrl(val string) modFunc {
	return modFunc(func(cfg *routeConfig) {
		cfg.TemplateUrl = val
	})
}
