package route

import "github.com/gopherjs/gopherjs/js"

type Params struct {
	*js.Object `ajs-service:"$routeParams"`
}
