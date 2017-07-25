package three

import "github.com/gopherjs/gopherjs/js"

type PerspectiveCameraPosition struct {
	Z int
}

type PerspectiveCamera struct {
	*js.Object

	Position Vector3 `js:"position"`
}

func NewPerspectiveCamera(fov, aspect, near, far float64) PerspectiveCamera {
	return PerspectiveCamera{Object: three.Get("PerspectiveCamera").New(fov, aspect, near, far)}
}

func (c PerspectiveCamera) Copy() PerspectiveCamera {
	return PerspectiveCamera{Object: c.Object.Call("copy")}
}

func (c PerspectiveCamera) SetFocalLength(focalLength float64) {
	c.Object.Call("setFocalLength", focalLength)
}

func (c PerspectiveCamera) GetFocalLength() float64 {
	return c.Object.Call("getFocalLength").Float()
}

func (c PerspectiveCamera) GetEffectiveFOV() float64 {
	return c.Object.Call("getEffectiveFOV").Float()
}

func (c PerspectiveCamera) GetFilmWidth() float64 {
	return c.Object.Call("getFilmWidth").Float()
}

func (c PerspectiveCamera) GetFilmHeight() float64 {
	return c.Object.Call("getFilmHeight").Float()
}

func (c PerspectiveCamera) SetViewOffset(fullWidth, fullHeight, x, y, width, height float64) {
	c.Object.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}

func (c PerspectiveCamera) ClearViewOffset() {
	c.Object.Call("clearViewOffset")
}

func (c PerspectiveCamera) UpdateProjectionMatrix() {
	c.Object.Call("updateProjectionMatrix")
}

func (c PerspectiveCamera) ToJSON(meta interface{}) interface{} {
	return c.Object.Call("toJSON", meta).Interface()
}
