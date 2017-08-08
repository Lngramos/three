package three

//go:generate go run object3d_method_generator/main.go -typeName AmbientLight -typeSlug ambient_light

import "github.com/gopherjs/gopherjs/js"

// AmbientLight - a light that gets emitted in a specific direction.
type AmbientLight struct {
	*js.Object

	Position *Vector3 `js:"position"`
}

func NewAmbientLight(color *Color, intensity float64) *AmbientLight {
	return &AmbientLight{
		Object: three.Get("AmbientLight").New(color, intensity),
	}
}
