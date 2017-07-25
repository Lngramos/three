package three

import (
	"github.com/gopherjs/gopherjs/js"
)

type Mesh struct {
	*js.Object

	Rotation struct {
		X float64 `js:"x"`
		Y float64 `js:"y"`
	} `js:"rotation"`
}

func NewMesh(geometry BoxGeometry, material MeshBasicMaterial) Mesh {
	return Mesh{
		Object: three.Get("Mesh").New(geometry.Object, material.Object),
	}
}

func (m Mesh) SetRotationFromAxisAngle(axis string, angle float64) {
	m.Call("setRotationFromAxisAngle", axis, angle)
}

func (m Mesh) RotateX() {
	m.Call("rotateX")
}
