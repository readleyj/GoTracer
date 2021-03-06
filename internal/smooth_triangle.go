package internal

import (
	"math"
	"math/rand"
	"time"
)

type SmoothTriangle struct {
	ID               int64
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
	Parent           Shape
	P1, P2, P3       Tuple
	N1, N2, N3       Tuple
	E1, E2           Tuple
	Normal           Tuple
	HasShadow        bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) *SmoothTriangle {
	e1, e2 := SubTuples(p2, p1), SubTuples(p3, p1)

	return &SmoothTriangle{
		ID:               rand.Int63(),
		Material:         NewDefaultMaterial(),
		Transform:        NewIdentity4(),
		Inverse:          NewIdentity4(),
		InverseTranspose: NewIdentity4(),
		Parent:           nil,
		P1:               p1,
		P2:               p2,
		P3:               p3,
		N1:               n1,
		N2:               n2,
		N3:               n3,
		E1:               e1,
		E2:               e2,
		Normal:           Normalize(Cross(e2, e1)),
		HasShadow:        true,
	}
}

func (tri *SmoothTriangle) GetID() int64 {
	return tri.ID
}

func (tri *SmoothTriangle) LocalIntersect(localRay Ray) Intersections {
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
	return NewIntersections(NewIntersectionUV(t, tri, u, v))
}

func (tri *SmoothTriangle) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return AddTuples(
		AddTuples(
			TupleScalarMultiply(tri.N2, hit.U),
			TupleScalarMultiply(tri.N3, hit.V),
		),
		TupleScalarMultiply(tri.N1, (1-hit.U-hit.V)),
	)
}

func (tri *SmoothTriangle) GetTransform() Matrix {
	return tri.Transform
}

func (tri *SmoothTriangle) SetTransform(transform Matrix) {
	tri.Transform = transform
	tri.Inverse = MatrixInverse(tri.Transform)
	tri.InverseTranspose = MatrixTranspose(tri.Inverse)
}

func (tri *SmoothTriangle) GetInverse() Matrix {
	return tri.Inverse
}

func (tri *SmoothTriangle) GetInverseTranspose() Matrix {
	return tri.InverseTranspose
}

func (tri *SmoothTriangle) GetMaterial() Material {
	return tri.Material
}

func (tri *SmoothTriangle) SetMaterial(material Material) {
	tri.Material = material
}

func (tri *SmoothTriangle) GetParent() Shape {
	return tri.Parent
}

func (tri *SmoothTriangle) SetParent(s Shape) {
	tri.Parent = s
}

func (tri *SmoothTriangle) CastsShadow() bool {
	return tri.HasShadow
}
