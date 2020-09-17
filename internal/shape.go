package internal

import "github.com/google/go-cmp/cmp"

type Shape interface {
	GetID() int64

	GetTransform() Matrix
	SetTransform(t Matrix)
	GetInverse() Matrix
	GetInverseTranspose() Matrix

	GetMaterial() Material
	SetMaterial(m Material)

	GetParent() Shape
	SetParent(s Shape)

	LocalIntersect(localRay Ray) Intersections
	LocalNormalAt(point Tuple, i Intersection) Tuple

	CastsShadow() bool
}

type TestShape struct {
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
	Parent           Shape
	SavedRay         Ray
	HasShadow        bool
}

func NewTestShape() *TestShape {
	return &TestShape{
		Material:         NewDefaultMaterial(),
		Transform:        NewIdentity4(),
		Inverse:          NewIdentity4(),
		InverseTranspose: NewIdentity4(),
		Parent:           nil,
		SavedRay:         Ray{},
		HasShadow:        true,
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
	t.Inverse = MatrixInverse(t.Transform)
	t.InverseTranspose = MatrixTranspose(t.Inverse)
}

func (t *TestShape) GetInverse() Matrix {
	return t.Inverse
}

func (t *TestShape) GetInverseTranspose() Matrix {
	return t.InverseTranspose
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

func (t *TestShape) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return NewVector(point.X, point.Y, point.Z)
}

func (t *TestShape) GetParent() Shape {
	return t.Parent
}

func (t *TestShape) SetParent(s Shape) {
	t.Parent = s
}

func (t *TestShape) CastsShadow() bool {
	return t.HasShadow
}

func Intersect(s Shape, ray Ray) Intersections {
	r := TransformRay(ray, s.GetInverse())
	return s.LocalIntersect(r)
}

func ShapesAreIdentical(s1, s2 Shape) bool {
	return s1.GetID() == s2.GetID()
}

func ShapeEquals(s1, s2 Shape) bool {
	return cmp.Equal(s1.GetMaterial(), s2.GetMaterial(), opt) && MatrixEquals(s1.GetTransform(), s2.GetTransform())
}

func ShapeHasParent(s Shape) bool {
	return s.GetParent() != nil
}

func WorldToObject(s Shape, point Tuple) Tuple {
	if ShapeHasParent(s) {
		point = WorldToObject(s.GetParent(), point)
	}

	return MatrixTupleMultiply(s.GetInverse(), point)
}

func NormalToWorld(s Shape, normal Tuple) Tuple {
	normal = MatrixTupleMultiply(s.GetInverseTranspose(), normal)
	normal.W = 0
	normal = Normalize(normal)

	if ShapeHasParent(s) {
		normal = NormalToWorld(s.GetParent(), normal)
	}

	return normal
}
