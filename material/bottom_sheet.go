package material

import "github.com/gopherjs/gopherjs/js"

type BottomSheet struct {
	*js.Object `ajs-service:"$mdBottomSheet"`
}

func (s *BottomSheet) Show(sheet *Sheet) {
	s.Call("show", sheet)
}

func (s *BottomSheet) Hide(val interface{}) {
	s.Call("hide", val)
}

func (s *BottomSheet) Cancel(val interface{}) {
	s.Call("cancel", val)
}
