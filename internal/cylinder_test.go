package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayMissesCylinder(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
	}{
		{NewPoint(1, 0, 0), NewVector(0, 1, 0)},
		{NewPoint(0, 0, 0), NewVector(0, 1, 0)},
		{NewPoint(0, 0, -5), NewVector(1, 1, 1)},
	}

	cyl := NewCylinder()

	for _, test := range testCases {
		direction := Normalize(test.direction)
		ray := NewRay(test.origin, direction)
		xs := cyl.LocalIntersect(ray)

		assert.Equal(t, 0, len(xs))
	}
}

func TestRayStrikesCylinder(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
		t0        float64
		t1        float64
	}{
		{NewPoint(1, 0, -5), NewVector(0, 0, 1), 5, 5},
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 4, 6},
		{NewPoint(0.5, 0, -5), NewVector(0.1, 1, 1), 6.80798, 7.08872},
	}

	cyl := NewCylinder()

	for _, test := range testCases {
		direction := Normalize(test.direction)
		ray := NewRay(test.origin, direction)
		xs := cyl.LocalIntersect(ray)

		assert.Equal(t, 2, len(xs))
		assert.InDelta(t, test.t0, xs[0].T, float64EqualityThreshold)
		assert.InDelta(t, test.t1, xs[1].T, float64EqualityThreshold)
	}
}

func TestNormalOnCylinder(t *testing.T) {
	testCases := []struct {
		point  Tuple
		normal Tuple
	}{
		{NewPoint(1, 0, 0), NewVector(1, 0, 0)},
		{NewPoint(0, 5, -1), NewVector(0, 0, -1)},
		{NewPoint(0, -2, 1), NewVector(0, 0, 1)},
		{NewPoint(-1, 1, 0), NewVector(-1, 0, 0)},
	}

	cyl := NewCylinder()

	for _, test := range testCases {
		n := cyl.LocalNormalAt(test.point)

		assert.True(t, TupleEquals(test.normal, n))
	}
}

func TestDefaultBoundsForCylinder(t *testing.T) {
	cyl := NewCylinder()

	assert.Equal(t, math.Inf(-1), cyl.Minimum)
	assert.Equal(t, math.Inf(1), cyl.Maximum)
}

func TestIntersectConstrainedCylinder(t *testing.T) {
	testCases := []struct {
		point     Tuple
		direction Tuple
		count     int
	}{
		{NewPoint(0, 1.5, 0), NewVector(0.1, 1, 0), 0},
		{NewPoint(0, 3, -5), NewVector(0, 0, 1), 0},
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 0},
		{NewPoint(0, 2, -5), NewVector(0, 0, 1), 0},
		{NewPoint(0, 1, -5), NewVector(0, 0, 1), 0},
		{NewPoint(0, 1.5, -2), NewVector(0, 0, 1), 2},
	}

	cyl := NewCylinder()
	cyl.Minimum = 1
	cyl.Maximum = 2

	for _, test := range testCases {
		direction := Normalize(test.direction)
		r := NewRay(test.point, direction)
		xs := cyl.LocalIntersect(r)

		assert.Equal(t, test.count, len(xs))
	}
}

func TestCylinderDefaultClosedValue(t *testing.T) {
	cyl := NewCylinder()

	assert.False(t, cyl.Closed)
}

func TestIntersectCapsOfClosedCylinder(t *testing.T) {
	testCases := []struct {
		point     Tuple
		direction Tuple
		count     int
	}{
		{NewPoint(0, 3, 0), NewVector(0, -1, 0), 2},
		{NewPoint(0, 3, -2), NewVector(0, -1, 2), 2},
		{NewPoint(0, 4, -2), NewVector(0, -1, 1), 2},
		{NewPoint(0, 0, -2), NewVector(0, 1, 2), 2},
		{NewPoint(0, -1, -2), NewVector(0, 1, 1), 2},
	}

	cyl := NewCylinder()
	cyl.Minimum = 1
	cyl.Maximum = 2
	cyl.Closed = true

	for _, test := range testCases {
		direction := Normalize(test.direction)
		r := NewRay(test.point, direction)
		xs := cyl.LocalIntersect(r)

		assert.Equal(t, test.count, len(xs))
	}

}

func TestCylinderNormalOnEndCaps(t *testing.T) {
	testCases := []struct {
		point  Tuple
		normal Tuple
	}{
		{NewPoint(0, 1, 0), NewVector(0, -1, 0)},
		{NewPoint(0.5, 1, 0), NewVector(0, -1, 0)},
		{NewPoint(0, 1, 0.5), NewVector(0, -1, 0)},
		{NewPoint(0, 2, 0), NewVector(0, 1, 0)},
		{NewPoint(0.5, 2, 0), NewVector(0, 1, 0)},
		{NewPoint(0, 2, 0.5), NewVector(0, 1, 0)},
	}

	cyl := NewCylinder()
	cyl.Minimum = 1
	cyl.Maximum = 2
	cyl.Closed = true

	for _, test := range testCases {
		n := cyl.LocalNormalAt(test.point)

		assert.True(t, TupleEquals(test.normal, n))
	}
}
