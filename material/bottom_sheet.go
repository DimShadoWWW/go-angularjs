package material

import "github.com/gopherjs/gopherjs/js"
import "github.com/AmandaCameron/go-angularjs"

type Sheet interface {
	BuildSheet() *js.Object
}

type CustomSheet struct {
	*js.Object

	Template    string            `js:"template"`
	TemplateURL string            `js:"templateUrl"`
	Locals      map[string]string `js:"locals"`

	Controller interface{}
}

func (sheet *CustomSheet) BuildSheet() *js.Object {
	if str, ok := sheet.Controller.(string); ok {
		sheet.Set("controller", str)
	} else {
		ctrler, err := angularjs.Annotate(sheet.Controller, angularjs.AnnotServices)
		if err != nil {
			panic(err)
		}

		sheet.Set("controller", ctrler)
	}

	return sheet.Object
}

func NewSheet() *CustomSheet {
	return &CustomSheet{
		Object: js.Global.Get("Object").New(),
	}
}

type BottomSheet struct {
	*js.Object `ajs-service:"$mdBottomSheet"`
}

func (s *BottomSheet) Show(sheet Sheet) {
	s.Call("show", sheet.BuildSheet())
}

func (s *BottomSheet) Hide(val interface{}) {
	s.Call("hide", val)
}

func (s *BottomSheet) Cancel(val interface{}) {
	s.Call("cancel", val)
}
