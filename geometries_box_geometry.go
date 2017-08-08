package three

//go:generate go run geometry_method_generator/main.go -geometryType BoxGeometry -geometrySlug box_geometry

import "github.com/gopherjs/gopherjs/js"

// BoxGeometry is the quadrilateral primitive geometry class. It is typically
// used for creating a cube or irregular quadrilateral of the dimensions
// provided with the 'width', 'height', and 'depth' constructor arguments.
type BoxGeometry struct {
	*js.Object

	Width          float64 `js:"width"`
	Height         float64 `js:"height"`
	Depth          float64 `js:"depth"`
	WidthSegments  float64 `js:"widthSegments"`
	HeightSegments float64 `js:"heightSegments"`
	DepthSegments  float64 `js:"depthSegments"`
}

// BoxGeometryParameters .
type BoxGeometryParameters struct {
	Width          float64
	Height         float64
	Depth          float64
	WidthSegments  float64
	HeightSegments float64
	DepthSegments  float64
}

// NewBoxGeometry creates a new BoxGeometry.
func NewBoxGeometry(params *BoxGeometryParameters) BoxGeometry {
	return BoxGeometry{
		Object: three.Get("BoxGeometry").New(
			params.Width,
			params.Height,
			params.Depth,
			params.WidthSegments,
			params.HeightSegments,
			params.DepthSegments,
		),
	}
}
