package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSmoothTriangle(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)
	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	assert.True(t, TupleEquals(p1, tri.P1))
	assert.True(t, TupleEquals(p2, tri.P2))
	assert.True(t, TupleEquals(p3, tri.P3))
	assert.True(t, TupleEquals(n1, tri.N1))
	assert.True(t, TupleEquals(n2, tri.N2))
	assert.True(t, TupleEquals(n3, tri.N3))
}

func TestSmoothTriangleIntersectionStoresUV(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)
	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	assert.InDelta(t, 0.45, xs[0].U, float64EqualityThreshold)
	assert.InDelta(t, 0.25, xs[0].V, float64EqualityThreshold)
}

func TestSmoothTriangleInterpolatesNormalWithUV(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)
	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	i := NewIntersectionUV(1, tri, 0.45, 0.25)
	n := NormalAt(tri, NewPoint(0, 0, 0), i)

	assert.True(t, TupleEquals(NewVector(-0.5547, 0.83205, 0), n))
}

func TestPrepareNormalOnSmoothTriangle(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)
	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	i := NewIntersectionUV(1, tri, 0.45, 0.25)
	r := NewRay(NewPoint(-0.2, 0.3, -2.0), NewVector(0, 0, 1))
	xs := NewIntersections(i)
	comps := PrepareComputations(i, r, xs)

	assert.True(t, TupleEquals(NewVector(-0.5547, 0.83205, 0), comps.NormalV))
}
