package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrecompIntersectionState(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(4, shape)
	comps := PrepareComputations(i, r, NewIntersections(i))

	assert.InDelta(t, i.T, comps.T, float64EqualityThreshold)
	assert.True(t, ShapeEquals(i.Object, comps.Object))
	assert.True(t, TupleEquals(NewPoint(0, 0, -1), comps.Point))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), comps.EyeV))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), comps.NormalV))
}

func TestPrecomputeReflectionVector(t *testing.T) {
	shape := NewPlane()
	r := NewRay(NewPoint(0, 1, -1), NewVector(0, -1/math.Sqrt(2), 1/math.Sqrt(2)))
	i := NewIntersection(math.Sqrt(2), shape)
	comps := PrepareComputations(i, r, NewIntersections(i))

	assert.True(t, comps.ReflectV.IsVector())
	assert.InDelta(t, 0.0, comps.ReflectV.X, float64EqualityThreshold)
	assert.InDelta(t, 1/math.Sqrt(2), comps.ReflectV.Y, float64EqualityThreshold)
	assert.InDelta(t, 1/math.Sqrt(2), comps.ReflectV.Z, float64EqualityThreshold)
}

func TestHitShouldOffsetPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	shape.SetTransform(Translate(0, 0, 1))
	i := NewIntersection(5, shape)
	comps := PrepareComputations(i, r, NewIntersections(i))

	assert.Less(t, comps.OverPoint.Z, -float64EqualityThreshold/2)
	assert.Greater(t, comps.Point.Z, comps.OverPoint.Z)
}

func TestUnderPointIsOffsetBelowSurface(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewGlassSphere()
	shape.SetTransform(Translate(0, 0, 1))
	i := NewIntersection(5, shape)
	xs := NewIntersections(i)
	comps := PrepareComputations(i, r, xs)

	assert.Greater(t, comps.UnderPoint.Z, -float64EqualityThreshold/2)
	assert.Less(t, comps.Point.Z, comps.UnderPoint.Z)
}
