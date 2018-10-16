package three

import "github.com/gopherjs/gopherjs/js"

// Face3 is a triangular face used in Geometry. These are created automatically for all standard geometry types, however if you are building a custom geometry you will have to create them manually.
type Face3 struct {
	*js.Object
}

// NewFace3 creates new Face3 object with defaults for normal, color and material.
func NewFace3(a, b, c float64) Face3 {
	return Face3{
		Object: three.Get("Face3").New(a, b, c),
	}
}
