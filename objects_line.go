package three

//go:generate go run object3d_method_generator/main.go -typeName Line -typeSlug line

import (
	"github.com/gopherjs/gopherjs/js"
)

// Line a continuous line.
type Line struct {
	*js.Object
}

// NewLine creates a new material. If Material is nil, three.js will assign a randomized material to the line o_O.
func NewLine(geom Geometry, material Material) *Line {
	return &Line{
		Object: three.Get("Line").New(geom, material),
	}
}
