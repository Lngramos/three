package three

import "github.com/gopherjs/gopherjs/js"

// Quaternion - represents a Quaternion.
type Quaternion struct {
	*js.Object
}

func NewQuaternion() Quaternion {
	return Quaternion{
		Object: three.Get("Quaternion").New(),
	}
}
