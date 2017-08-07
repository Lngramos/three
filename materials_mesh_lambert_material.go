package three

//go:generate go run cmd/object3d_method_generator.go -typeName MeshLambertMaterial -typeSlug mesh_lambert_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)

type MeshLambertMaterial struct {
	*js.Object
}

func NewMeshLambertMaterial(params *MaterialParameters) *MeshLambertMaterial {
	return &MeshLambertMaterial{
		Object: three.Get("MeshLambertMaterial").New(params.Object),
	}
}

func (m MeshLambertMaterial) SetValues(params *MaterialParameters) {
	m.Call("setValues", structs.Map(params))
}

func (m MeshLambertMaterial) getObject() *js.Object {
	return m.Object
}
