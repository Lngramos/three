package three

import "github.com/gopherjs/gopherjs/js"

// BoxGeometry is the quadrilateral primitive geometry class. It is typically
// used for creating a cube or irregular quadrilateral of the dimensions
// provided with the 'width', 'height', and 'depth' constructor arguments.
type BoxGeometry struct {
	*js.Object
}

// BoxGeometryParameters .
type BoxGeometryParameters struct {
	*js.Object

	Width          float64 `js:"width"`
	Height         float64 `js:"height"`
	Depth          float64 `js:"depth"`
	WidthSegments  float64 `js:"widthSegments"`
	HeightSegments float64 `js:"heightSegments"`
	DepthSegments  float64 `js:"depthSegments"`
}

// NewBoxGeometry creates a new BoxGeometry.
func NewBoxGeometry(width, height, depth float64) BoxGeometry {
	return BoxGeometry{
		Object: three.Get("BoxGeometry").New(width, height, depth),
	}
}
