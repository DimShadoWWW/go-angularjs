package material

import "github.com/gopherjs/gopherjs/js"
import "github.com/AmandaCameron/go-angularjs"

type DialogService struct {
	*js.Object `ajs-service:"$mdDialog"`
}

func (s *DialogService) Alert() *PresetDialog {
	return &PresetDialog{s.Call("alert")}
}

func (s *DialogService) Show(dlg Dialog) {
	s.Call("show", dlg.BuildDialog())
}

func (s *DialogService) Hide(dlg Dialog) {
	if dlg != nil {
		s.Call("hide", dlg.BuildDialog())
	} else {
		s.Call("hide")
	}
}

type CustomDialog struct {
	*js.Object

	Controller interface{}

	Template            string            `js:"template"`
	TemplateURL         string            `js:"templateUrl"`
	Locals              map[string]string `js:"locals"`
	Backdrop            bool              `js:"hasBackdrop"`
	EscapeToClose       bool              `js:"escapeToClose"`
	ClickOutsideToClose bool              `js:"clickOutsideToClose"`
	FocusOnOpen         bool              `js:"focusOnOpen"`
	DisableParentScroll bool              `js:"disableParentScroll"`
}

func NewDialog() *CustomDialog {
	return &CustomDialog{
		Object: js.Global.Get("Object").New(),
	}
}

func (dlg *CustomDialog) BuildDialog() *js.Object {
	if str, ok := dlg.Controller.(string); ok {
		dlg.Set("controller", str)
	} else {
		controller, err := angularjs.Annotate(dlg.Controller, angularjs.AnnotServices)
		if err != nil {
			panic(err)
		}

		dlg.Set("controller", controller)
	}

	return dlg.Object
}

type Dialog interface {
	BuildDialog() *js.Object
}

type PresetDialog struct {
	*js.Object
}

func (dlg *PresetDialog) BuildDialog() *js.Object {
	return dlg.Object
}

func (dlg *PresetDialog) Theme(val string) {
	dlg.Call("theme", val)
}

func (dlg *PresetDialog) Title(val string) {
	dlg.Call("title", val)
}

func (dlg *PresetDialog) Content(val string) {
	dlg.Call("content", val)
}

func (dlg *PresetDialog) Cancel(val string) {
	dlg.Call("cancel", val)
}

func (dlg *PresetDialog) Ok(val string) {
	dlg.Call("ok", val)
}
