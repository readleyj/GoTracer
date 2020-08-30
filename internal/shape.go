package internal

import "github.com/google/go-cmp/cmp"

type Shape interface {
	GetID() int64
	GetTransform() Matrix
	SetTransform(t Matrix)

	GetMaterial() Material
	SetMaterial(m Material)

	GetParent() *Group
	SetParent(g *Group)

	LocalIntersect(localRay Ray) Intersections
	LocalNormalAt(point Tuple, i Intersection) Tuple
}

type TestShape struct {
	Material  Material
	Transform Matrix
	Parent    *Group
	SavedRay  Ray
}

func NewTestShape() *TestShape {
	return &TestShape{
		NewDefaultMaterial(),
		NewIdentity4(),
		nil,
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

func (t *TestShape) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return NewVector(point.X, point.Y, point.Z)
}

func (t *TestShape) GetParent() *Group {
	return t.Parent
}

func (t *TestShape) SetParent(g *Group) {
	t.Parent = g
}

func Intersect(s Shape, ray Ray) Intersections {
	r := TransformRay(ray, MatrixInverse(s.GetTransform()))
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

	return MatrixTupleMultiply(MatrixInverse(s.GetTransform()), point)
}

func NormalToWorld(s Shape, normal Tuple) Tuple {
	normal = MatrixTupleMultiply(MatrixTranspose(MatrixInverse(s.GetTransform())), normal)
	normal.W = 0
	normal = Normalize(normal)

	if ShapeHasParent(s) {
		normal = NormalToWorld(s.GetParent(), normal)
	}

	return normal
}
