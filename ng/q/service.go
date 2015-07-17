package q

import "github.com/gopherjs/gopherjs/js"

type DeferPromise struct {
	*js.Object
}

func (dp *DeferPromise) Promise() *js.Object {
	return dp.Get("promise")
}

func (dp *DeferPromise) Resolve(val interface{}) {
	dp.Call("resolve", val)
}

func (dp *DeferPromise) Notify(val interface{}) {
	dp.Call("notify", val)
}

func (dp *DeferPromise) Reject(val interface{}) {
	dp.Call("reject", val)
}

type Service struct {
	*js.Object `ajs-service:"$q"`
}

func (s *Service) Defer() *DeferPromise {
	return &DeferPromise{s.Call("defer")}
}
