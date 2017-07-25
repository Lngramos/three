package three

import "github.com/gopherjs/gopherjs/js"

// Matrix4 - represents a Matrix4.
type Matrix4 struct {
	*js.Object
}

func NewMatrix4(x, y, z float64) Matrix4 {
	return Matrix4{
		Object: three.Get("Matrix4").New(x, y, z),
	}
}
