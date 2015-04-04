package http

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$http"`
}

func (s *Service) Get(url string, callback func(data *Response, status int)) {
	future := s.Call("get", url)

	future.Call("success", func(data *Response, status int, headers *js.Object, config *js.Object) {
		callback(data, status)
	})

	future.Call("error", func(data *Response, status int, headers *js.Object, config *js.Object) {
		callback(data, status)
	})
}
