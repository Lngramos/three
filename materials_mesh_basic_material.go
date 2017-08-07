package three

//go:generate go run cmd/object3d_method_generator.go -typeName MeshBasicMaterial -typeSlug mesh_basic_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)

type MeshBasicMaterial struct {
	*js.Object
}

func NewMeshBasicMaterial(params *MaterialParameters) *MeshBasicMaterial {
	return &MeshBasicMaterial{
		Object: three.Get("MeshBasicMaterial").New(params.Object),
	}
}

func (m MeshBasicMaterial) SetValues(params *MaterialParameters) {
	m.Call("setValues", structs.Map(params))
}

func (m MeshBasicMaterial) getObject() *js.Object {
	return m.Object
}
