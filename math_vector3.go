package three

import "github.com/gopherjs/gopherjs/js"

// Vector3 - represents a Vector3.
type Vector3 struct {
	*js.Object
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		Object: three.Get("Vector3").New(x, y, z),
	}
}

func (v Vector3) Set(x, y, z float64) Vector3 {
	v.Call("set", x, y, z)
	return v
}
