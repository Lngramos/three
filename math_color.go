package three

import "github.com/gopherjs/gopherjs/js"

// Color - represents a color.
type Color struct {
	*js.Object
}

func NewColor(r, g, b float64) *Color {
	return &Color{
		Object: three.Get("Color").New(r, g, b),
	}
}
