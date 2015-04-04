package http

import (
	"encoding/json"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

type Response struct {
	*js.Object
}

func (r *Response) Json(val interface{}) error {
	// FIXME: This is a massive cludge.
	return json.NewDecoder(strings.NewReader(js.Global.Get("angular").Call("toJson", r).String())).Decode(val)
}
