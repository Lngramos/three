package three

//go:generate go run geometry_method_generator/main.go -geometryType CircleGeometry -geometrySlug circle_geometry

import "github.com/gopherjs/gopherjs/js"

type CircleGeometry struct {
	*js.Object

	Radius      float64 `js:"radius"`
	Segments    float64 `js:"segments"`
	ThetaStart  float64 `js:"thetaStart"`
	ThetaLength float64 `js:"thetaLength"`
}

type CircleGeometryParameters struct {
	Radius      float64
	Segments    float64
	ThetaStart  float64
	ThetaLength float64
}

// NewBoxGeometry creates a new BoxGeometry.
func NewCircleGeometry(params CircleGeometryParameters) CircleGeometry {
	return CircleGeometry{
		Object: three.Get("CircleGeometry").New(
			params.Radius,
			params.Segments,
			params.ThetaStart,
			params.ThetaLength,
		),
	}
}
