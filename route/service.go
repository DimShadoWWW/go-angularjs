package route

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$route"`
}

func (s *Service) Reload() {
	s.Call("reload")
}

func (s *Service) UpdateParams(params map[string]string) {
	s.Call("updateParams", params)
}
