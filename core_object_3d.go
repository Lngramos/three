package three

import "github.com/gopherjs/gopherjs/js"

type Object3D interface {
	ApplyMatrix(matrix *Matrix4)
	Add(m Object3D)
	ToJSON() interface{}
	getInternalObject() *js.Object
}

// OnBeforeRender()
// OnAfterRender()
// ApplyMatrix(matrix float64)
// ApplyQuaternion(q Quaternion)
// SetRotationFromAxisAngle(axis Vector3, angle float64)
// SetRotationFromEuler(euler Euler)
// SetRotationFromMatrix(m Matrix4)
// SetRotationFromQuaternion(q Quaternion)
// RotateOnAxis()
// RotateX()
// RotateY()
// RotateZ()
// TranslateOnAxis()
// TranslateX()
// TranslateY()
// TranslateZ()
// LocalToWorld(vector Vector3)
// WorldToLocal()
// LookAt()
// Add(object Object3D)
// Remove(object Object3D)
// GetObjectById(id string)
// GetObjectByName(name string)
// GetObjectByProperty(name, value string)
// GetWorldPosition(optionalTarget Object3D)
// GetWorldQuaternion()
// GetWorldRotation()
// GetWorldScale()
// GetWorldDirection()
// Raycast()
// UpdateMatrix()
// UpdateMatrixWorld(force bool)
// ToJSON(meta interface{})
// Clone(recursive bool)
// Copy(source Object3D, recursive bool)
