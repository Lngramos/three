package three

//go:generate go run object3d_method_generator/main.go -typeName Scene -typeSlug scene

import (
	"github.com/gopherjs/gopherjs/js"
)

// Scene - Allows you to set up what and where is to be rendered by three.js. This is where you place objects, lights and cameras.
type Scene struct {
	*js.Object
}

// NewScene - Create a new Scene object.
func NewScene() *Scene {
	return &Scene{Object: three.Get("Scene").New()}
}
