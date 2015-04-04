package anchor_scroll

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$anchorScroll"`
}

func (s *Service) Scroll(hash string) {
	if hash != "" {
		s.Invoke(hash)
	} else {
		s.Invoke()
	}
}
