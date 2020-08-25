package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRay(t *testing.T) {
	origin := NewPoint(1, 2, 3)
	direction := NewVector(4, 5, 6)
	r := NewRay(origin, direction)

	assert.True(t, origin.Equals(r.Origin))
	assert.True(t, direction.Equals(r.Direction))
}

func TestComputePointFromDistance(t *testing.T) {
	r := NewRay(NewPoint(2, 3, 4), NewVector(1, 0, 0))

	testCases := []struct {
		target   Tuple
		distance float64
	}{
		{NewPoint(2, 3, 4), 0.0},
		{NewPoint(3, 3, 4), 1.0},
		{NewPoint(1, 3, 4), -1.0},
		{NewPoint(4.5, 3, 4), 2.5},
	}

	for _, test := range testCases {
		assert.True(t, test.target.Equals(Position(r, test.distance)))
	}
}

func TestRayIntersectSphereAtTwoPoints(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 4.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 6.0, xs[1].T, float64EqualityThreshold)
}

func TestRayIsTangentToSphere(t *testing.T) {
	r := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 5.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 5.0, xs[1].T, float64EqualityThreshold)
}

func TestRayMissesSphere(t *testing.T) {
	r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 0, len(xs))
}

func TestRayOriginatesInSphere(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, -1.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 1.0, xs[1].T, float64EqualityThreshold)
}

func TestSphereIsBehindRay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, -6.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, -4.0, xs[1].T, float64EqualityThreshold)
}

func TestTranslateRay(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	m := Translate(3, 4, 5)
	r2 := TransformRay(r, m)

	assert.True(t, r2.Origin.Equals(NewPoint(4, 6, 8)))
	assert.True(t, r2.Direction.Equals(NewVector(0, 1, 0)))
}

func TestScaleRay(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVector(0, 1, 0))
	m := Scale(2, 3, 4)
	r2 := TransformRay(r, m)

	assert.True(t, r2.Origin.Equals(NewPoint(2, 6, 12)))
	assert.True(t, r2.Direction.Equals(NewVector(0, 3, 0)))
}
