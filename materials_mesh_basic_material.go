package three

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)

type MeshBasicMaterialParameters struct {
	*js.Object
	Color Color `js:"color"`
}

type MeshBasicMaterial struct {
	*js.Object
}

func NewMeshBasicMaterial(params MeshBasicMaterialParameters) MeshBasicMaterial {
	return MeshBasicMaterial{
		Object: three.Get("MeshBasicMaterial").New(params),
	}
}

func (m MeshBasicMaterial) SetValues(params MeshBasicMaterialParameters) {
	m.Call("setValues", structs.Map(params))
}
