package three

import "github.com/gopherjs/gopherjs/js"

// Euler - represents a Euler.
type Euler struct {
	*js.Object

	X     float64 `js:"x"`
	Y     float64 `js:"y"`
	Z     float64 `js:"z"`
	Order float64 `js:"order"`
}

func NewEuler(x, y, z, order float64) Euler {
	return Euler{
		Object: three.Get("Euler").New(x, y, z, order),
	}
}
