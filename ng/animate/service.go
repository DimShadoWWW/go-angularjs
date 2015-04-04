package animate

import "github.com/gopherjs/gopherjs/js"

type Service struct {
	*js.Object `ajs-service:"$animate"`
}

func (s *Service) Enter(elem *JQueryElement, parent *JQueryElement, after *JQueryElement, styles *js.Object) AnimationPromise {
	return AnimationPromise{s.Call("enter", elem, parent, after, styles)}
}

func (s *Service) Leave(elem *JQueryElement, styles *js.Object) AnimationPromise {
	return AnimationPromise{s.Call("leave", elem, styles)}
}

func (s *Service) Move(elem *JQueryElement, parent *JQueryElement, after *JQueryElement, styles *js.Object) AnimationPromise {
	return AnimationPromise{s.Call("move", elem, parent, after, styles)}
}

func (s *Service) AddClass(elem *JQueryElement, clsName string, options *js.Object) AnimationPromise {
	return AnimationPromise{s.Call("addClass", elem, clsName, options)}
}

func (s *Service) RemoveClass(elem *JQueryElement, clsName string, options *js.Object) AnimationPromise {
	return AnimationPromise{s.Call("removeClass", elem, clsName, options)}
}

func (s *Service) SetClass(elem *JQueryElement, add, remove string, options *js.Object) AnimationPromise {
	return AnimationPromise{s.Call("setClass", elem, add, remove, options)}
}
