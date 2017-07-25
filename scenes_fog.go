package three

import "github.com/gopherjs/gopherjs/js"

type Fog struct {
	*js.Object

	Color string  `js:"color"`
	Near  float64 `js:"near"`
	Far   float64 `js:"far"`
}

func NewFog(color Color, near float64, far float64) Fog {
	fog := three.Get("Fog").New(color.Object, near, far)

	return Fog{
		Object: fog,
		Color:  fog.Get("color").String(),
		Near:   fog.Get("near").Float(),
		Far:    fog.Get("far").Float(),
	}
}
