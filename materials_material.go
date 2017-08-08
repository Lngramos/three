package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// Side defines which side of faces will be rendered - front, back or both. Default is FrontSide.
type Side float64

const (
	FrontSide  Side = 0
	BackSide   Side = 1
	DoubleSide Side = 2
)

type Shading float64

const (
	// SmoothShading is the default and linearly interpolates color between vertices.
	SmoothShading Shading = 0

	// FlatShading uses the color of the first vertex for every pixel in a face.
	FlatShading Shading = 1
)

type MaterialParameters struct {
	*js.Object

	Color   *Color  `js:"color"`
	Shading Shading `js:"shading"`
	Side    Side    `js:"side"`
}

func NewMaterialParameters() *MaterialParameters {
	return &MaterialParameters{
		Object: js.Global.Get("Object").New(),
	}
}

type Material interface {
	OnBeforeCompile()
	SetValues(values MaterialParameters)
	ToJSON(meta interface{}) interface{}
	Clone()
	Copy(source Object3D)
	Dispose()

	getInternalObject() *js.Object
}
