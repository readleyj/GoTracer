package internal

import (
	"math"
	"math/rand"
	"time"
)

type Triangle struct {
	ID         int64
	Material   Material
	Transform  Matrix
	Parent     Shape
	P1, P2, P3 Tuple
	E1, E2     Tuple
	Normal     Tuple
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewTriangle(p1, p2, p3 Tuple) *Triangle {
	e1, e2 := SubTuples(p2, p1), SubTuples(p3, p1)

	return &Triangle{
		ID:        rand.Int63(),
		Material:  NewDefaultMaterial(),
		Transform: NewIdentity4(),
		Parent:    nil,
		P1:        p1,
		P2:        p2,
		P3:        p3,
		E1:        e1,
		E2:        e2,
		Normal:    Normalize(Cross(e2, e1)),
	}
}

func (tri *Triangle) GetID() int64 {
	return tri.ID
}

func (tri *Triangle) LocalIntersect(localRay Ray) Intersections {
	dirCrossE2 := Cross(localRay.Direction, tri.E2)
	det := Dot(tri.E1, dirCrossE2)

	if math.Abs(det) < float64EqualityThreshold {
		return Intersections{}
	}

	f := 1.0 / det

	p1ToOrigin := SubTuples(localRay.Origin, tri.P1)
	u := f * Dot(p1ToOrigin, dirCrossE2)

	if u < 0 || u > 1 {
		return Intersections{}
	}

	originCrossE1 := Cross(p1ToOrigin, tri.E1)
	v := f * Dot(localRay.Direction, originCrossE1)

	if v < 0 || (u+v) > 1 {
		return Intersections{}
	}

	t := f * Dot(tri.E2, originCrossE1)
	return NewIntersections(NewIntersection(t, tri))
}

func (tri *Triangle) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return tri.Normal
}

func (tri *Triangle) GetTransform() Matrix {
	return tri.Transform
}

func (tri *Triangle) SetTransform(transform Matrix) {
	tri.Transform = transform
}

func (tri *Triangle) GetMaterial() Material {
	return tri.Material
}

func (tri *Triangle) SetMaterial(material Material) {
	tri.Material = material
}

func (tri *Triangle) GetParent() Shape {
	return tri.Parent
}

func (tri *Triangle) SetParent(s Shape) {
	tri.Parent = s
}
