package three

//go:generate go run cmd/object3d_method_generator.go -typeName Mesh -typeSlug mesh

import (
	"github.com/gopherjs/gopherjs/js"
)

type Mesh struct {
	*js.Object

	Rotation struct {
		*js.Object

		X float64 `js:"x"`
		Y float64 `js:"y"`
	} `js:"rotation"`
}

// NewMesh aaa bbb ccc
func NewMesh(geometry *BoxGeometry, material Object3D) *Mesh {
	return &Mesh{
		Object: three.Get("Mesh").New(geometry.Object, material.getInternalObject()),
	}
}

func (m Mesh) SetRotationFromAxisAngle(axis string, angle float64) {
	m.Call("setRotationFromAxisAngle", axis, angle)
}

func (m Mesh) RotateX() {
	m.Call("rotateX")
}
