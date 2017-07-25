package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Scene - Allows you to set up what and where is to be rendered by three.js. This is where you place objects, lights and cameras.
type Scene struct {
	*js.Object
}

// NewScene - Create a new Scene object.
func NewScene() Scene {
	return Scene{Object: three.Get("Scene").New()}
}

func (s Scene) Add(m Mesh) {
	s.Object.Call("add", m.Object)
}

func (s Scene) Copy() Scene {
	return Scene{Object: s.Object.Call("copy")}
}

func (s Scene) ToJSON() interface{} {
	return s.Object.Call("toJSON").Interface()
}
