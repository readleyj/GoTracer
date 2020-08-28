package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRayIntersectsCube(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
		t1        float64
		t2        float64
	}{
		{NewPoint(5, 0.5, 0), NewVector(-1, 0, 0), 4, 6},
		{NewPoint(-5, 0.5, 0), NewVector(1, 0, 0), 4, 6},
		{NewPoint(0.5, 5, 0), NewVector(0, -1, 0), 4, 6},
		{NewPoint(0.5, -5, 0), NewVector(0, 1, 0), 4, 6},
		{NewPoint(0.5, 0, 5), NewVector(0, 0, -1), 4, 6},
		{NewPoint(0.5, 0, -5), NewVector(0, 0, 1), 4, 6},
		{NewPoint(0, 0.5, 0), NewVector(0, 0, 1), -1, 1},
	}

	c := NewCube()

	for _, test := range testCases {
		ray := NewRay(test.origin, test.direction)
		xs := c.LocalIntersect(ray)

		assert.Equal(t, 2, len(xs))
		assert.InDelta(t, test.t1, xs[0].T, float64EqualityThreshold)
		assert.InDelta(t, test.t2, xs[1].T, float64EqualityThreshold)
	}

}

func TestRayMissesCube(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
	}{
		{NewPoint(-2, 0, 0), NewVector(0.2673, 0.5345, 0.8018)},
		{NewPoint(0, -2, 0), NewVector(0.8018, 0.2673, 0.5345)},
		{NewPoint(0, 0, -2), NewVector(0.5345, 0.8018, 0.2673)},
		{NewPoint(2, 0, 2), NewVector(0, 0, -1)},
		{NewPoint(0, 2, 2), NewVector(0, -1, 0)},
		{NewPoint(2, 2, 0), NewVector(-1, 0, 0)},
	}

	c := NewCube()

	for _, test := range testCases {
		ray := NewRay(test.origin, test.direction)
		xs := c.LocalIntersect(ray)

		assert.Equal(t, 0, len(xs))
	}
}

func TestNormalOnCubeSurface(t *testing.T) {
	testCases := []struct {
		point  Tuple
		normal Tuple
	}{
		{NewPoint(1, 0.5, -0.8), NewVector(1, 0, 0)},
		{NewPoint(-1, -0.2, 0.9), NewVector(-1, 0, 0)},
		{NewPoint(-0.4, 1, -0.1), NewVector(0, 1, 0)},
		{NewPoint(0.3, -1, -0.7), NewVector(0, -1, 0)},
		{NewPoint(-0.6, 0.3, 1), NewVector(0, 0, 1)},
		{NewPoint(0.4, 0.4, -1), NewVector(0, 0, -1)},
		{NewPoint(1, 1, 1), NewVector(1, 0, 0)},
		{NewPoint(-1, -1, -1), NewVector(-1, 0, 0)},
	}

	c := NewCube()

	for _, test := range testCases {
		p := test.point
		normal := c.LocalNormalAt(p)

		assert.True(t, TupleEquals(test.normal, normal))
	}
}
