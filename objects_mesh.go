package three

//go:generate go run object3d_method_generator/main.go -typeName Mesh -typeSlug mesh

import (
	"github.com/gopherjs/gopherjs/js"
)

type Mesh struct {
	*js.Object

	Position *Vector3 `js:"position"`
	Rotation *Euler   `js:"rotation"`
}

func NewMesh(geometry Geometry, material Material) *Mesh {
	return &Mesh{
		Object: three.Get("Mesh").New(geometry.getInternalObject(), material.getInternalObject()),
	}
}

func (m Mesh) SetRotationFromAxisAngle(axis string, angle float64) {
	m.Call("setRotationFromAxisAngle", axis, angle)
}

func (m Mesh) RotateX() {
	m.Call("rotateX")
}
