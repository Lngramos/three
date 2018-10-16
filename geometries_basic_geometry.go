package three

//go:generate go run geometry_method_generator/main.go -geometryType BasicGeometry -geometrySlug basic_geometry

import "github.com/gopherjs/gopherjs/js"

// BasicGeometry is the basic primitive geometry class. It's a counterpart of three.js' core/Geometry object.
type BasicGeometry struct {
	*js.Object
}

type BasicGeometryParams struct {
}

// NewBasicGeometry creates a new Basic Geometry.
func NewBasicGeometry(params BasicGeometryParams) BasicGeometry {
	return BasicGeometry{
		Object: three.Get("Geometry").New(),
	}
}

// AddVertice adds new vertice to the geometry, specified by its coordinates.
func (bg *BasicGeometry) AddVertice(x, y, z float64) {
	vec := NewVector3(x, y, z)
	bg.Get("vertices").Call("push", vec)
}

// AddFace adds new Face3 (triangle) to the geometry, specified by its vertice indicies.
func (bg *BasicGeometry) AddFace(a, b, c int) {
	face := NewFace3(float64(a), float64(b), float64(c))
	bg.Get("faces").Call("push", face)
}
