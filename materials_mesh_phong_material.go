package three

//go:generate go run cmd/object3d_method_generator.go -typeName MeshPhongMaterial -typeSlug mesh_phong_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)

type MeshPhongMaterial struct {
	*js.Object
}

func NewMeshPhongMaterial(params *MaterialParameters) *MeshPhongMaterial {
	return &MeshPhongMaterial{
		Object: three.Get("MeshPhongMaterial").New(params.Object),
	}
}

func (m MeshPhongMaterial) SetValues(params *MaterialParameters) {
	m.Call("setValues", structs.Map(params))
}

func (m MeshPhongMaterial) getObject() *js.Object {
	return m.Object
}
