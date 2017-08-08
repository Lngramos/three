package three

import "github.com/gopherjs/gopherjs/js"

type Geometry interface {
	ApplyMatrix(matrix *Matrix4)
	RotateX()
	RotateY()
	RotateZ()
	Translate()
	Scale()
	LookAt()
	FromBufferGeometry(geometry Geometry)
	Center()
	// Normalize()
	ComputeFaceNormals()
	ComputeVertexNormals(areaWeighted bool)
	ComputeFlatVertexNormals()
	ComputeMorphNormals()
	ComputeLineDistances()
	ComputeBoundingBox()
	ComputeBoundingSphere()
	Merge(geometry Geometry, matrix Matrix4, materialIndexOffset float64)
	MergeMesh(mesh Mesh)
	MergeVertices()
	SortFacesByMaterialIndex()
	ToJSON() interface{}
	// Clone()
	// Copy(source Object3D)
	Dispose()

	getInternalObject() *js.Object
}
