package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/lngramos/three"
)

var (
	scene    three.Scene
	camera   three.PerspectiveCamera
	renderer three.WebGLRenderer
	mesh     three.Mesh
)

// var camera = new THREE.PerspectiveCamera( 70, window.innerWidth / window.innerHeight, 1, 1000 );
// camera.position.z = 400;

// var scene = new THREE.Scene();

// var geometry = new THREE.BoxBufferGeometry( 200, 200, 200 );
// var material = new THREE.MeshBasicMaterial( { color: "red" } );

// var mesh = new THREE.Mesh( geometry, material );
// scene.add( mesh );

// var renderer = new THREE.WebGLRenderer();
// renderer.setPixelRatio( window.devicePixelRatio );
// renderer.setSize( window.innerWidth, window.innerHeight );
// document.body.appendChild( renderer.domElement );

// renderer.render( scene, camera );

func main() {
	document := js.Global.Get("document")
	windowWidth := js.Global.Get("innerWidth").Float()
	windowHeight := js.Global.Get("innerHeight").Float()
	devicePixelRatio := js.Global.Get("devicePixelRatio").Float()

	camera = three.NewPerspectiveCamera(70, windowWidth/windowHeight, 1, 1000)
	camera.Position.Set(0, 0, 400)

	scene = three.NewScene()
	// camera.Position.Z = 5

	renderer = three.NewWebGLRenderer()
	renderer.SetPixelRatio(devicePixelRatio)
	renderer.SetSize(windowWidth, windowHeight, true)
	document.Get("body").Call("appendChild", renderer.Get("domElement"))

	// Create cube
	geometry := three.NewBoxGeometry(200, 200, 200)
	material := three.NewMeshBasicMaterial(three.MeshBasicMaterialParameters{Color: three.NewColor(255, 0, 0)})
	mesh = three.NewMesh(geometry, material)
	scene.Add(mesh)

	renderer.Render(scene, camera)
	// animate()
}

func animate() {
	js.Global.Call("requestAnimationFrame", animate)

	// cube.SetRotationFromAxisAngle("x", cube.Rotation.X+0.01)
	// cube.SetRotationFromAxisAngle("y", cube.Rotation.Y+0.01)
	// cube.Rotation.X += 0.01
	// cube.Rotation.Y += 0.02
	// cube.RotateX()

	renderer.Render(scene, camera)
}
