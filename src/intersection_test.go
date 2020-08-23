package internal

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestIntersectionEncapsulatesObjectAndT(t *testing.T) {
	s := MakeSphere()
	i := MakeIntersection(3.5, &s)

	assert.InDelta(t, 3.5, i.T, float64EqualityThreshold)
	assert.Equal(t, s.ID, (*i.Object).ID)
}

func TestAggregatingIntersections(t *testing.T) {
	s := MakeSphere()
	i1 := MakeIntersection(1, &s)
	i2 := MakeIntersection(2, &s)
	xs := MakeIntersections(i1, i2)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 1.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 2.0, xs[1].T, float64EqualityThreshold)
}

func TestIntersectSetsObjectOnIntersection(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s.ID, (*xs[0].Object).ID)
	assert.Equal(t, s.ID, (*xs[1].Object).ID)
}

func TestHitWhenAllIntersectionsHavePositiveT(t *testing.T) {
	s := MakeSphere()
	i1 := MakeIntersection(1, &s)
	i2 := MakeIntersection(2, &s)
	xs := MakeIntersections(i2, i1)
	i := Hit(xs)

	assert.InDelta(t, i1.T, i.T, float64EqualityThreshold)
	assert.Equal(t, i1.Object.ID, i.Object.ID)
}

func TestHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	s := MakeSphere()
	i1 := MakeIntersection(-1, &s)
	i2 := MakeIntersection(1, &s)
	xs := MakeIntersections(i2, i1)
	i := Hit(xs)

	assert.InDelta(t, i2.T, i.T, float64EqualityThreshold)
	assert.Equal(t, i2.Object.ID, i.Object.ID)
}

func TestHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := MakeSphere()
	i1 := MakeIntersection(-2, &s)
	i2 := MakeIntersection(-1, &s)
	xs := MakeIntersections(i2, i1)
	i := Hit(xs)
	empty := Intersection{}

	assert.True(t, cmp.Equal(empty.Object, i.Object))
}

func TestHitIsLowestNonnegativeIntersection(t *testing.T) {
	s := MakeSphere()
	i1 := MakeIntersection(5, &s)
	i2 := MakeIntersection(7, &s)
	i3 := MakeIntersection(-3, &s)
	i4 := MakeIntersection(2, &s)
	xs := MakeIntersections(i1, i2, i3, i4)
	i := Hit(xs)

	assert.InDelta(t, i4.T, i.T, float64EqualityThreshold)
	assert.Equal(t, i4.Object.ID, i.Object.ID)
}

func TestIntersectScaledSphereWithRay(t *testing.T) {
	r := Ray{Point(0, 0, -5), Vector(0, 0, 1)}
	s := MakeSphere()
	s.SetTransform(Scale(2, 2, 2))
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 3.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 7.0, xs[1].T, float64EqualityThreshold)
}
