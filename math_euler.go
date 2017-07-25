package three

import "github.com/gopherjs/gopherjs/js"

// Euler - represents a Euler.
type Euler struct {
	*js.Object
}

func NewEuler(x, y, z, order float64) Euler {
	return Euler{
		Object: three.Get("Euler").New(x, y, z, order),
	}
}
