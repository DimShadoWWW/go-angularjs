package route

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$router"`
}

func (r Service) When(path string, mods ...modFunc) {
	r.addPath(path, mods)
}

func (r Service) Otherwise(mods ...modFunc) {
	//AddPath(r.Module, r.Controller, "/", mods...)
	r.addPath("/", mods)
}

func (r Service) Recognise(url string) bool {
	return r.Call("recognize", url).Bool()
}

func (r Service) Navigate(url string) {
	r.Call("navigate", url)
}

func (r Service) addPath(path string, mods []modFunc) {
	value := &routeConfig{
		Object: js.Global.Get("Object").New(),
	}

	value.Path = path

	for _, mod := range mods {
		mod(value)
	}

	r.Call("config", value)
}
