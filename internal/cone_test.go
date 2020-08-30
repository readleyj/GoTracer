package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayIntersectCone(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
		t0        float64
		t1        float64
	}{
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 5, 5},
		{NewPoint(0, 0, -5), NewVector(1, 1, 1), 8.66025, 8.66025},
		{NewPoint(1, 1, -5), NewVector(-0.5, -1, 1), 4.55006, 49.44994},
	}

	cone := NewCone()

	for _, test := range testCases {
		direction := Normalize(test.direction)
		ray := NewRay(test.origin, direction)
		xs := cone.LocalIntersect(ray)

		assert.Equal(t, 2, len(xs))
		assert.InDelta(t, test.t0, xs[0].T, float64EqualityThreshold)
		assert.InDelta(t, test.t1, xs[1].T, float64EqualityThreshold)
	}
}

func TestIntersectConeWithParallelRay(t *testing.T) {
	cone := NewCone()

	direction := Normalize(NewVector(0, 1, 1))
	ray := NewRay(NewPoint(0, 0, -1), direction)
	xs := cone.LocalIntersect(ray)

	assert.Equal(t, 1, len(xs))
	assert.InDelta(t, 0.35355, xs[0].T, float64EqualityThreshold)
}

func TestConeIntersectEndCaps(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
		count     int
	}{
		{NewPoint(0, 0, -5), NewVector(0, 1, 0), 0},
		{NewPoint(0, 0, -0.25), NewVector(0, 1, 1), 2},
		{NewPoint(0, 0, -0.25), NewVector(0, 1, 0), 4},
	}

	cone := NewCone()
	cone.Minimum = -0.5
	cone.Maximum = 0.5
	cone.Closed = true

	for _, test := range testCases {
		direction := Normalize(test.direction)
		ray := NewRay(test.origin, direction)
		xs := cone.LocalIntersect(ray)

		assert.Equal(t, test.count, len(xs))
	}
}

func TestNormalOnCone(t *testing.T) {
	testCases := []struct {
		point  Tuple
		normal Tuple
	}{
		{NewPoint(0, 0, 0), NewVector(0, 0, 0)},
		{NewPoint(1, 1, 1), NewVector(1, -math.Sqrt(2), 1)},
		{NewPoint(-1, -1, 0), NewVector(-1, 1, 0)},
	}

	cone := NewCone()

	for _, test := range testCases {
		n := cone.LocalNormalAt(test.point, Intersection{})

		assert.True(t, TupleEquals(test.normal, n))
	}
}
