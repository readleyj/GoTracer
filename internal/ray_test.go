package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRay(t *testing.T) {
	origin := Point(1, 2, 3)
	direction := Vector(4, 5, 6)
	r := Ray{origin, direction}

	assert.True(t, origin.Equals(r.Origin))
	assert.True(t, direction.Equals(r.Direction))
}

func TestComputePointFromDistance(t *testing.T) {
	r := Ray{Point(2, 3, 4), Vector(1, 0, 0)}

	testCases := []struct {
		target   Tuple
		distance float64
	}{
		{Point(2, 3, 4), 0.0},
		{Point(3, 3, 4), 1.0},
		{Point(1, 3, 4), -1.0},
		{Point(4.5, 3, 4), 2.5},
	}

	for _, test := range testCases {
		assert.True(t, test.target.Equals(Position(r, test.distance)))
	}
}

func TestRayIntersectSphereAtTwoPoints(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 4.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 6.0, xs[1].T, float64EqualityThreshold)
}

func TestRayIsTangentToSphere(t *testing.T) {
	r := Ray{Point(0, 1, -5), Vector(0, 0, 1)}
	s := MakeSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 5.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 5.0, xs[1].T, float64EqualityThreshold)
}

func TestRayMissesSphere(t *testing.T) {
	r := Ray{Point(0, 2, -5), Vector(0, 0, 1)}
	s := MakeSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 0, len(xs))
}

func TestRayOriginatesInSphere(t *testing.T) {
	r := Ray{Point(0, 0, 0), Vector(0, 0, 1)}
	s := MakeSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, -1.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 1.0, xs[1].T, float64EqualityThreshold)
}

func TestSphereIsBehindRay(t *testing.T) {
	r := Ray{Point(0, 0, 5), Vector(0, 0, 1)}
	s := MakeSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, -6.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, -4.0, xs[1].T, float64EqualityThreshold)
}

func TestTranslateRay(t *testing.T) {
	r := Ray{Point(1, 2, 3), Vector(0, 1, 0)}
	m := Translate(3, 4, 5)
	r2 := TransformRay(r, m)

	assert.True(t, r2.Origin.Equals(Point(4, 6, 8)))
	assert.True(t, r2.Direction.Equals(Vector(0, 1, 0)))
}

func TestScaleRay(t *testing.T) {
	r := Ray{Point(1, 2, 3), Vector(0, 1, 0)}
	m := Scale(2, 3, 4)
	r2 := TransformRay(r, m)

	assert.True(t, r2.Origin.Equals(Point(2, 6, 12)))
	assert.True(t, r2.Direction.Equals(Vector(0, 3, 0)))
}
