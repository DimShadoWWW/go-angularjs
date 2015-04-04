package provide

import "github.com/gopherjs/gopherjs/js"

type Provider struct {
	*js.Object `ajs-provider:"$provide"`
}
