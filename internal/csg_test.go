package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCSG(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	c := NewCSG(CSGUnion, s1, s2)

	assert.Equal(t, CSGUnion, c.Operation)
	assert.True(t, ShapeEquals(s1, c.Left))
	assert.True(t, ShapeEquals(s2, c.Right))
	assert.True(t, ShapeEquals(c, s1.GetParent()))
	assert.True(t, ShapeEquals(c, s2.GetParent()))
}

func TestRuleForCSGUnion(t *testing.T) {
	testCases := []struct {
		op     CSGOperation
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{CSGUnion, true, true, true, false},
		{CSGUnion, true, true, false, true},
		{CSGUnion, true, false, true, false},
		{CSGUnion, true, false, false, true},
		{CSGUnion, false, true, true, false},
		{CSGUnion, false, true, false, false},
		{CSGUnion, false, false, true, true},
		{CSGUnion, false, false, false, true},
	}

	for _, test := range testCases {
		result := IntersectionAllowed(test.op, test.lhit, test.inl, test.inr)
		assert.Equal(t, test.result, result)
	}
}

func TestRuleForCSGIntersection(t *testing.T) {
	testCases := []struct {
		op     CSGOperation
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{CSGIntersect, true, true, true, true},
		{CSGIntersect, true, true, false, false},
		{CSGIntersect, true, false, true, true},
		{CSGIntersect, true, false, false, false},
		{CSGIntersect, false, true, true, true},
		{CSGIntersect, false, true, false, true},
		{CSGIntersect, false, false, true, false},
		{CSGIntersect, false, false, false, false},
	}

	for _, test := range testCases {
		result := IntersectionAllowed(test.op, test.lhit, test.inl, test.inr)
		assert.Equal(t, test.result, result)
	}
}

func TestRuleForCSGDifference(t *testing.T) {
	testCases := []struct {
		op     CSGOperation
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{CSGDifference, true, true, true, false},
		{CSGDifference, true, true, false, true},
		{CSGDifference, true, false, true, false},
		{CSGDifference, true, false, false, true},
		{CSGDifference, false, true, true, true},
		{CSGDifference, false, true, false, true},
		{CSGDifference, false, false, true, false},
		{CSGDifference, false, false, false, false},
	}

	for _, test := range testCases {
		result := IntersectionAllowed(test.op, test.lhit, test.inl, test.inr)
		assert.Equal(t, test.result, result)
	}
}

func TestFilterIntersections(t *testing.T) {
	testCases := []struct {
		op CSGOperation
		x0 int
		x1 int
	}{
		{CSGUnion, 0, 3},
		{CSGIntersect, 1, 2},
		{CSGDifference, 0, 1},
	}

	s1 := NewSphere()
	s2 := NewCube()

	intersections := []Intersection{
		NewIntersection(1, s1),
		NewIntersection(2, s2),
		NewIntersection(3, s1),
		NewIntersection(4, s2),
	}
	xs := NewIntersections(intersections...)

	for _, test := range testCases {
		c := NewCSG(test.op, s1, s2)

		result := FilterIntersections(c, xs)
		assert.Equal(t, 2, len(result))
		assert.True(t, IntersectEquals(xs[test.x0], result[0]))
		assert.True(t, IntersectEquals(xs[test.x1], result[1]))
	}
}

func TestRayMissesCSG(t *testing.T) {
	c := NewCSG(CSGUnion, NewSphere(), NewCube())
	r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
	xs := c.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestRayHitsCSG(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(Translate(0, 0, 0.5))

	c := NewCSG(CSGUnion, s1, s2)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := c.LocalIntersect(r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 4.0, xs[0].T, float64EqualityThreshold)
	assert.True(t, Includes(xs[0].Object, s1))
	assert.InDelta(t, 6.5, xs[1].T, float64EqualityThreshold)
	assert.True(t, Includes(xs[1].Object, s2))
}
