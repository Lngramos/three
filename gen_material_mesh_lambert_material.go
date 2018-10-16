package three
// Code generated by go generate; DO NOT EDIT.
//
// using the following cmd:
// material_method_generator -materialName MeshLambertMaterial -materialSlug mesh_lambert_material

import (
	"github.com/fatih/structs"
	"github.com/gopherjs/gopherjs/js"
)
	
func (m MeshLambertMaterial) OnBeforeCompile() {
	m.Call("onBeforeCompile")
}

func (m MeshLambertMaterial) SetValues(values MaterialParameters) {
	m.Call("setValues", structs.Map(values))
}

func (m MeshLambertMaterial) ToJSON(meta interface{}) interface{} {
	return m.Call("toJSON", meta)
}

func (m MeshLambertMaterial) Clone() {
	m.Call("clone")
}

func (m MeshLambertMaterial) Copy(source Object3D) {
	m.Call("copy", source)
}

func (m MeshLambertMaterial) Dispose() {
	m.Call("dispose")
}

func (m MeshLambertMaterial) getInternalObject() *js.Object {
	return m.Object
}

