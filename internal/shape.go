package internal

import "github.com/google/go-cmp/cmp"

type Shape interface {
	GetID() int64
	GetTransform() Matrix
	SetTransform(t Matrix)

	GetMaterial() Material
	SetMaterial(m Material)

	LocalIntersect(localRay Ray) Intersections
	LocalNormalAt(point Tuple) Tuple
}

type TestShape struct {
	Material  Material
	Transform Matrix
	SavedRay  Ray
}

func NewTestShape() *TestShape {
	return &TestShape{
		NewDefaultMaterial(),
		NewIdentity4(),
		Ray{},
	}
}

func (t *TestShape) GetID() int64 {
	return 0
}

func (t *TestShape) GetTransform() Matrix {
	return t.Transform
}

func (t *TestShape) SetTransform(transform Matrix) {
	t.Transform = transform
}

func (t *TestShape) GetMaterial() Material {
	return t.Material
}

func (t *TestShape) SetMaterial(material Material) {
	t.Material = material
}

func (t *TestShape) LocalIntersect(localRay Ray) Intersections {
	t.SavedRay = localRay
	return Intersections{}
}

func (t *TestShape) LocalNormalAt(point Tuple) Tuple {
	return NewVector(point.X, point.Y, point.Z)
}

func Intersect(s Shape, ray Ray) Intersections {
	r := TransformRay(ray, MatrixInverse(s.GetTransform()))
	return s.LocalIntersect(r)
}

func ShapeEquals(s1, s2 Shape) bool {
	return cmp.Equal(s1.GetMaterial(), s2.GetMaterial(), opt) && MatrixEquals(s1.GetTransform(), s2.GetTransform())
}
