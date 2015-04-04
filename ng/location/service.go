package location

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-controller:"$location"`
}

func (s *Service) GetUrl() string {
	return s.Call("url").String()
}

func (s *Service) SetUrl(url string) {
	s.Call("url", url)
}

func (s *Service) GetHash() string {
	return s.Call("hash").String()
}

func (s *Service) SetHash(hash string) {
	s.Call("hash", hash)
}

func (s *Service) GetPath() string {
	return s.Call("path").String()
}

func (s *Service) SetPath(path string) {
	s.Call("path", path)
}

//func (s *Service) GetSearch() (ret map[string]string) {e
//	s.Call("search").Interface(ret)
//}

func (s *Service) SetSearch(search map[string]string) {
	s.Call("search", search)
}

func (s *Service) GetSearchPart(name string) string {
	return s.Call("search", name).String()
}

func (s *Service) SetSearchPart(name, value string) {
	s.Call("search", name, value)
}

func (s *Service) GetState() *js.Object {
	return s.Call("state")
}

func (s *Service) SetState(state *js.Object) {
	s.Call("state", state)
}
