package material

import "github.com/gopherjs/gopherjs/js"

type DialogService struct {
	*js.Object `ajs-service:"$mdDialog"`
}

func (s *DialogService) Alert() *PresetDialog {
	return &PresetDialog{s.Call("alert")}
}

func (s *DialogService) Show(dlg Dialog) {
	s.Call("show", dlg.Build())
}

func (s *DialogService) Hide(dlg Dialog) {
	s.Call("hide", dlg.Build())
}

type CustomDialog struct {
	*js.Object

	Template            string            `js:"template"`
	Controller          string            `js:"controller"`
	Locals              map[string]string `js:"locals"`
	Backdrop            bool              `js:"hasBackdrop"`
	EscapeToClose       bool              `js:"escapeToClose"`
	ClickOutsideToClose bool              `js:"clickOutsideToClose"`
	FocusOnOpen         bool              `js:"focusOnOpen"`
	DisableParentScroll bool              `js:"disableParentScroll"`
}

func (dlg *CustomDialog) Build() *js.Object {
	return dlg.Object
}

type Dialog interface {
	Build() *js.Object
}

type PresetDialog struct {
	*js.Object
}

func (dlg *PresetDialog) Build() *js.Object {
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
