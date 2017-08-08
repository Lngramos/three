package three

//go:generate go run object3d_method_generator/main.go -typeName DirectionalLight -typeSlug directional_light

import "github.com/gopherjs/gopherjs/js"

// DirectionalLight - a light that gets emitted in a specific direction.
type DirectionalLight struct {
	*js.Object

	Position *Vector3 `js:"position"`
}

func NewDirectionalLight(color *Color, intensity float64) *DirectionalLight {
	return &DirectionalLight{
		Object: three.Get("DirectionalLight").New(color, intensity),
	}
}
