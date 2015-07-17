package material

import "github.com/gopherjs/gopherjs/js"

type IconProvider struct {
	*js.Object `ajs-provider:"$mdIconProvider"`
}

func (pro *IconProvider) Load(name, path string, size int) {
	pro.Call("iconSet", name, path, size)
}

func (pro *IconProvider) LoadDefault(path string, size int) {
	pro.Call("defaultIconSet", path, size)
}
