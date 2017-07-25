package three

import (
	"github.com/gopherjs/gopherjs/js"
)

// The WebGLRenderer displays your beautifully crafted scenes using WebGL.
type WebGLRenderer struct {
	*js.Object
}

// WebGLRenderer creates an WebGLRenderer instance.
func NewWebGLRenderer() WebGLRenderer {
	return WebGLRenderer{
		Object: three.Get("WebGLRenderer").New(),
	}
}

// WebGLRenderer creates an WebGLRenderer instance.
func (r WebGLRenderer) SetSize(w float64, h float64, updateStyle bool) {
	r.Call("setSize", w, h, updateStyle)
}

func (r WebGLRenderer) SetPixelRatio(ratio float64) {
	r.Call("setPixelRatio", ratio)
}

func (r WebGLRenderer) Render(scene Scene, camera PerspectiveCamera) {
	r.Call("render", scene.Object, camera.Object)
}
