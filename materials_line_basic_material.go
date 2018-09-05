package three

//go:generate go run material_method_generator/main.go -materialName LineBasicMaterial -materialSlug line_basic_material

import (
	"github.com/gopherjs/gopherjs/js"
)

type LineBasicMaterial struct {
	*js.Object
}

func NewLineBasicMaterial(params *MaterialParameters) *LineBasicMaterial {
	return &LineBasicMaterial{
		Object: three.Get("LineBasicMaterial").New(params.Object),
	}
}
