package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTriangle(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	tri := NewTriangle(p1, p2, p3)

	assert.True(t, TupleEquals(p1, tri.P1))
	assert.True(t, TupleEquals(p2, tri.P2))
	assert.True(t, TupleEquals(p3, tri.P3))
	assert.True(t, TupleEquals(NewVector(-1, -1, 0), tri.E1))
	assert.True(t, TupleEquals(NewVector(1, -1, 0), tri.E2))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), tri.Normal))
}

func TestFindNormalOnTriangle(t *testing.T) {
	tri := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	n1 := tri.LocalNormalAt(NewPoint(0, 0.5, 0), Intersection{})
	n2 := tri.LocalNormalAt(NewPoint(-0.5, 0.75, 0), Intersection{})
	n3 := tri.LocalNormalAt(NewPoint(0.5, 0.25, 0), Intersection{})

	assert.True(t, TupleEquals(n1, tri.Normal))
	assert.True(t, TupleEquals(n2, tri.Normal))
	assert.True(t, TupleEquals(n3, tri.Normal))
}

func TestIntersectRayParallelToTriangle(t *testing.T) {
	tri := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	r := NewRay(NewPoint(0, -1, -2), NewVector(0, 1, 0))
	xs := tri.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayMissesEdgeP13(t *testing.T) {
	tri := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	r := NewRay(NewPoint(1, 1, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayMissesEdgeP12(t *testing.T) {
	tri := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	r := NewRay(NewPoint(-1, 1, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayMissesEdgeP23(t *testing.T) {
	tri := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	r := NewRay(NewPoint(0, -1, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayStrikesTriangle(t *testing.T) {
	tri := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	r := NewRay(NewPoint(0, 0.5, -2), NewVector(0, 0, 1))
	xs := tri.LocalIntersect(r)

	assert.Equal(t, 1, len(xs))
	assert.InDelta(t, 2.0, xs[0].T, float64EqualityThreshold)
}
