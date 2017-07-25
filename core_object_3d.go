package three

import "github.com/gopherjs/gopherjs/js"

type Object3D struct {
	*js.Object

	Rotation struct {
		X float64 `js:"x"`
		Y float64 `js:"y"`
	} `js:"rotation"`
}

func (o Object3D) OnBeforeRender() {
	o.Call("onBeforeRender")
}

func (o Object3D) OnAfterRender() {
	o.Call("onAfterRender")
}

func (o Object3D) ApplyMatrix(matrix float64) {
	o.Call("applyMatrix", matrix)
}

func (o Object3D) ApplyQuaternion(q Quaternion) {
	o.Call("applyQuaternion")
}

func (o Object3D) SetRotationFromAxisAngle(axis Vector3, angle float64) {
	o.Call("setRotationFromAxisAngle")
}

func (o Object3D) SetRotationFromEuler(euler Euler) {
	o.Call("setRotationFromEuler")
}

func (o Object3D) SetRotationFromMatrix(m Matrix4) {
	o.Call("setRotationFromMatrix")
}

func (o Object3D) SetRotationFromQuaternion(q Quaternion) {
	o.Call("setRotationFromQuaternion", q)
}

func (o Object3D) RotateOnAxis() {
	o.Call("rotateOnAxis")
}

func (o Object3D) RotateX() {
	o.Call("rotateX")
}

func (o Object3D) RotateY() {
	o.Call("rotateY")
}

func (o Object3D) RotateZ() {
	o.Call("rotateZ")
}

func (o Object3D) TranslateOnAxis() {
	o.Call("translateOnAxis")
}

func (o Object3D) TranslateX() {
	o.Call("translateX")
}

func (o Object3D) TranslateY() {
	o.Call("translateY")
}

func (o Object3D) TranslateZ() {
	o.Call("translateZ")
}

func (o Object3D) LocalToWorld(vector Vector3) {
	o.Call("localToWorld")
}

func (o Object3D) WorldToLocal() {
	o.Call("worldToLocal")
}

func (o Object3D) LookAt() {
	o.Call("lookAt")
}

func (o Object3D) Add(object Object3D) {
	o.Call("add", object)
}

func (o Object3D) Remove(object Object3D) {
	o.Call("remove", object)
}

func (o Object3D) GetObjectById(id string) {
	o.Call("getObjectById", id)
}

func (o Object3D) GetObjectByName(name string) {
	o.Call("getObjectByName", name)
}

func (o Object3D) GetObjectByProperty(name, value string) {
	o.Call("getObjectByProperty", name, value)
}

func (o Object3D) GetWorldPosition(optionalTarget Object3D) {
	o.Call("getWorldPosition", optionalTarget)
}

func (o Object3D) GetWorldQuaternion() {
	o.Call("getWorldQuaternion")
}

func (o Object3D) GetWorldRotation() {
	o.Call("getWorldRotation")
}

func (o Object3D) GetWorldScale() {
	o.Call("getWorldScale")
}

func (o Object3D) GetWorldDirection() {
	o.Call("getWorldDirection")
}

func (o Object3D) Raycast() {
	o.Call("raycast")
}

func (o Object3D) UpdateMatrix() {
	o.Call("updateMatrix")
}

func (o Object3D) UpdateMatrixWorld(force bool) {
	o.Call("updateMatrixWorld", force)
}

func (o Object3D) ToJSON(meta interface{}) {
	o.Call("toJSON", meta)
}

func (o Object3D) Clone(recursive bool) {
	o.Call("clone")
}

func (o Object3D) Copy(source Object3D, recursive bool) {
	o.Call("copy")
}
