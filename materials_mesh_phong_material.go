package three

//go:generate go run material_method_generator/main.go -materialName MeshPhongMaterial -materialSlug mesh_phong_material

import (
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
