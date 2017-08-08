package three

//go:generate go run material_method_generator/main.go -materialName MeshLambertMaterial -materialSlug mesh_lambert_material

import (
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
