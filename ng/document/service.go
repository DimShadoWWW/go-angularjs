package document

import "github.com/gopherjs/gopherjs/js"
import "github.com/AmandaCameron/go-angularjs"

type Service struct {
	*js.Object `ajs-service:"$document"`
	*angularjs.JQueryElement
}

func (s *Service) Init() {
	s.JQueryElement = &angularjs.JQueryElement{s.Object}
}

func (s *Service) SetTitle(val string) {
	s.Index(0).Set("title", val)
}

func (s *Service) GetTitle() string {
	return s.Index(0).Get("title").String()
}
