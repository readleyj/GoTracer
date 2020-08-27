package internal

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestIntersectionEncapsulatesObjectAndT(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)

	assert.InDelta(t, 3.5, i.T, float64EqualityThreshold)
	assert.Equal(t, s.ID, i.Object.GetID())
}

func TestAggregatingIntersections(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2)

	assert.Equal(t, 2, len(xs))
	assert.InDelta(t, 1.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 2.0, xs[1].T, float64EqualityThreshold)
}

func TestIntersectSetsObjectOnIntersection(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	xs := Intersect(s, r)

	assert.Equal(t, 2, len(xs))
	assert.Equal(t, s.ID, xs[0].Object.GetID())
	assert.Equal(t, s.ID, xs[1].Object.GetID())
}

func TestHitWhenAllIntersectionsHavePositiveT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := NewIntersections(i2, i1)
	i := Hit(xs)

	assert.InDelta(t, i1.T, i.T, float64EqualityThreshold)
	assert.Equal(t, i1.Object.GetID(), i.Object.GetID())
}

func TestHitWhenSomeIntersectionsHaveNegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := NewIntersections(i2, i1)
	i := Hit(xs)

	assert.InDelta(t, i2.T, i.T, float64EqualityThreshold)
	assert.Equal(t, i2.Object.GetID(), i.Object.GetID())
}

func TestHitWhenAllIntersectionsHaveNegativeT(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := NewIntersections(i2, i1)
	i := Hit(xs)
	empty := Intersection{}

	assert.True(t, cmp.Equal(empty.Object, i.Object))
}

func TestHitIsLowestNonnegativeIntersection(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := NewIntersections(i1, i2, i3, i4)
	i := Hit(xs)

	assert.InDelta(t, i4.T, i.T, float64EqualityThreshold)
	assert.Equal(t, i4.Object.GetID(), i.Object.GetID())
}

func TestHitIntersectionOnOutside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(4, shape)
	comps := PrepareComputations(i, r)

	assert.False(t, comps.Inside)
}

func TestHitIntersectionOnInside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(1, shape)
	comps := PrepareComputations(i, r)

	assert.True(t, TupleEquals(NewPoint(0, 0, 1), comps.Point))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), comps.EyeV))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), comps.NormalV))
	assert.True(t, comps.Inside)
}

func TestShadeIntersection(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := w.Objects[0]
	i := NewIntersection(4, shape)
	comps := PrepareComputations(i, r)
	c := ShadeHit(w, comps)

	assert.InDelta(t, 0.38066, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.47583, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.28550, c.B, float64EqualityThreshold)
}

func TestShadeIntersectionFromInside(t *testing.T) {
	w := NewDefaultWorld()
	w.Lights[0] = NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1))
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := w.Objects[1]
	i := NewIntersection(0.5, shape)
	comps := PrepareComputations(i, r)
	c := ShadeHit(w, comps)

	assert.InDelta(t, 0.90498, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.90498, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.90498, c.B, float64EqualityThreshold)
}

func TestColorWhenRayMisses(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	c := ColorAt(w, r)

	assert.InDelta(t, 0.0, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.0, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.0, c.B, float64EqualityThreshold)
}

func TestColorWhenRayHits(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	c := ColorAt(w, r)

	assert.InDelta(t, 0.38066, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.47583, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.28550, c.B, float64EqualityThreshold)
}

func TestColorWithIntersectionBehindRay(t *testing.T) {
	w := NewDefaultWorld()

	outer := w.Objects[0]
	outerMaterial := outer.GetMaterial()
	outerMaterial.Ambient = 1
	outer.SetMaterial(outerMaterial)

	inner := w.Objects[1]
	innerMaterial := inner.GetMaterial()
	innerMaterial.Ambient = 1
	inner.SetMaterial(innerMaterial)

	r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))
	innerColor := inner.GetMaterial().Color
	c := ColorAt(w, r)

	assert.InDelta(t, innerColor.R, c.R, float64EqualityThreshold)
	assert.InDelta(t, innerColor.G, c.G, float64EqualityThreshold)
	assert.InDelta(t, innerColor.B, c.B, float64EqualityThreshold)
}

func TestShadeHitGivenIntersectionInShadow(t *testing.T) {
	w := NewWorld()
	w.Lights = append(w.Lights, NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)))
	s1 := NewSphere()
	w.Objects = append(w.Objects, s1)
	s2 := NewSphere()
	s2.SetTransform(Translate(0, 0, 10))
	w.Objects = append(w.Objects, s2)
	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	i := NewIntersection(4, s2)
	comps := PrepareComputations(i, r)
	c := ShadeHit(w, comps)

	assert.InDelta(t, 0.1, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.1, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.1, c.B, float64EqualityThreshold)
}

func TestHitShouldOffsetPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	shape.SetTransform(Translate(0, 0, 1))
	i := NewIntersection(5, shape)
	comps := PrepareComputations(i, r)

	assert.Less(t, comps.OverPoint.Z, -float64EqualityThreshold/2)
	assert.Greater(t, comps.Point.Z, comps.OverPoint.Z)
}
