package window

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$window"`
}
