package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaneNormalIsConstantEverywhere(t *testing.T) {
	p := NewPlane()
	n1 := p.LocalNormalAt(NewPoint(0, 0, 0), Intersection{})
	n2 := p.LocalNormalAt(NewPoint(10, 0, -10), Intersection{})
	n3 := p.LocalNormalAt(NewPoint(-5, 0, 150), Intersection{})

	assert.True(t, TupleEquals(NewVector(0, 1, 0), n1))
	assert.True(t, TupleEquals(NewVector(0, 1, 0), n2))
	assert.True(t, TupleEquals(NewVector(0, 1, 0), n3))
}

func TestPlaneIntersectRayParallel(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1))
	xs := p.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestPlaneIntersectRayCoplanar(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	xs := p.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayIntersectsPlaneFromAbove(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0))
	xs := p.LocalIntersect(r)

	assert.Equal(t, 1, len(xs))
	assert.InDelta(t, 1.0, xs[0].T, float64EqualityThreshold)
	assert.True(t, ShapeEquals(p, xs[0].Object))
}

func TestRayIntersectsPlaneFromBelow(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0))
	xs := p.LocalIntersect(r)

	assert.Equal(t, 1, len(xs))
	assert.InDelta(t, 1.0, xs[0].T, float64EqualityThreshold)
	assert.True(t, ShapeEquals(p, xs[0].Object))
}
