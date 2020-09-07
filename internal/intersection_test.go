package internal

import (
	"math"
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
	comps := PrepareComputations(i, r, NewIntersections(i))

	assert.False(t, comps.Inside)
}

func TestHitIntersectionOnInside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(1, shape)
	comps := PrepareComputations(i, r, NewIntersections(i))

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
	comps := PrepareComputations(i, r, NewIntersections(i))
	c := ShadeHit(w, comps, RecursionDepth)

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
	comps := PrepareComputations(i, r, NewIntersections(i))
	c := ShadeHit(w, comps, RecursionDepth)

	assert.InDelta(t, 0.90498, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.90498, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.90498, c.B, float64EqualityThreshold)
}

func TestColorWhenRayMisses(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	c := ColorAt(w, r, RecursionDepth)

	assert.InDelta(t, 0.0, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.0, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.0, c.B, float64EqualityThreshold)
}

func TestColorWhenRayHits(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	c := ColorAt(w, r, RecursionDepth)

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
	c := ColorAt(w, r, RecursionDepth)

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
	comps := PrepareComputations(i, r, NewIntersections(i))
	c := ShadeHit(w, comps, RecursionDepth)

	assert.InDelta(t, 0.1, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.1, c.G, float64EqualityThreshold)
	assert.InDelta(t, 0.1, c.B, float64EqualityThreshold)
}

func TestReflectedColorForNonreflectiveMaterial(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))

	shape := w.Objects[1]
	shapeMaterial := shape.GetMaterial()
	shape.SetMaterial(shapeMaterial)

	i := NewIntersection(1, shape)
	comps := PrepareComputations(i, r, NewIntersections(i))
	color := ReflectedColor(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0, 0, 0), color))
}

func TestReflectedColorForReflectiveMaterial(t *testing.T) {
	w := NewDefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.SetTransform(Translate(0, -1, 0))
	w.Objects = append(w.Objects, shape)

	r := NewRay(NewPoint(0.0, 0.0, -3.0), NewVector(0.0, -1/math.Sqrt(2), 1/math.Sqrt(2)))
	i := NewIntersection(math.Sqrt(2), shape)
	comps := PrepareComputations(i, r, NewIntersections(i))
	color := ReflectedColor(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0.19032, 0.23790, 0.14274), color))
}

func TestShadeHitWithReflectiveMaterial(t *testing.T) {
	w := NewDefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.SetTransform(Translate(0, -1, 0))
	w.Objects = append(w.Objects, shape)

	r := NewRay(NewPoint(0.0, 0.0, -3.0), NewVector(0.0, -1/math.Sqrt(2), 1/math.Sqrt(2)))
	i := NewIntersection(math.Sqrt(2), shape)
	comps := PrepareComputations(i, r, NewIntersections(i))
	color := ShadeHit(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0.87677, 0.92436, 0.82918), color))
}

func TestColorAtWithMutuallyReflectiveSurface(t *testing.T) {
	w := NewWorld()
	w.Lights = append(w.Lights, NewPointLight(NewPoint(0, 0, 0), NewColor(1, 1, 1)))

	lower := NewPlane()
	lowerMaterial := lower.GetMaterial()
	lowerMaterial.Reflective = 1
	lower.SetMaterial(lowerMaterial)
	lower.SetTransform(Translate(0, -1, 0))
	w.Objects = append(w.Objects, lower)

	upper := NewPlane()
	upperMaterial := upper.GetMaterial()
	upperMaterial.Reflective = 1
	upper.SetMaterial(upperMaterial)
	upper.SetTransform(Translate(0, 1, 0))
	w.Objects = append(w.Objects, upper)

	ray := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))
	ColorAt(w, ray, RecursionDepth)

	r := recover()
	assert.Nil(t, r)
}

func TestReflectedColorAtMaximumRecursiveDepth(t *testing.T) {
	w := NewDefaultWorld()
	shape := NewPlane()
	shape.Material.Reflective = 0.5
	shape.SetTransform(Translate(0, -1, 0))
	w.Objects = append(w.Objects, shape)

	r := NewRay(NewPoint(0.0, 0.0, -3.0), NewVector(0.0, -1/math.Sqrt(2), 1/math.Sqrt(2)))
	i := NewIntersection(math.Sqrt(2), shape)
	comps := PrepareComputations(i, r, NewIntersections(i))
	color := ReflectedColor(w, comps, 0)

	assert.True(t, ColorEquals(black, color))
}

func TestFindRefractiveIndicesAtIntersections(t *testing.T) {
	testCases := []struct {
		n1 float64
		n2 float64
	}{
		{1.0, 1.5},
		{1.5, 2.0},
		{2.0, 2.5},
		{2.5, 2.5},
		{2.5, 1.5},
		{1.5, 1.0},
	}

	A := NewGlassSphere()
	A.SetTransform(Scale(2, 2, 2))
	A.Material.RefractiveIndex = 1.5

	B := NewGlassSphere()
	B.SetTransform(Translate(0, 0, -0.25))
	B.Material.RefractiveIndex = 2.0

	C := NewGlassSphere()
	C.SetTransform(Translate(0, 0, 0.25))
	C.Material.RefractiveIndex = 2.5

	r := NewRay(NewPoint(0, 0, -4), NewVector(0, 0, 1))
	intersections := []Intersection{
		NewIntersection(2, A),
		NewIntersection(2.75, B),
		NewIntersection(3.25, C),
		NewIntersection(4.75, B),
		NewIntersection(5.25, C),
		NewIntersection(6, A),
	}
	xs := NewIntersections(intersections...)

	for index, test := range testCases {
		comps := PrepareComputations(xs[index], r, xs)
		assert.InDelta(t, test.n1, comps.N1, float64EqualityThreshold)
		assert.InDelta(t, test.n2, comps.N2, float64EqualityThreshold)
	}
}

func TestRefractedColorWithOpaqueSurface(t *testing.T) {
	w := NewDefaultWorld()
	shape := w.Objects[0]
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	intersections := []Intersection{
		NewIntersection(4, shape),
		NewIntersection(6, shape),
	}

	xs := NewIntersections(intersections...)
	comps := PrepareComputations(xs[0], r, xs)
	c := RefractedColor(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0, 0, 0), c))
}

func TestRefractedColorAtMaximumRecursiveDepth(t *testing.T) {
	w := NewDefaultWorld()
	shape := w.Objects[0]
	shapeMaterial := shape.GetMaterial()
	shapeMaterial.Transparency = 1.0
	shapeMaterial.RefractiveIndex = 1.5
	shape.SetMaterial(shapeMaterial)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	intersections := []Intersection{
		NewIntersection(4, shape),
		NewIntersection(6, shape),
	}

	xs := NewIntersections(intersections...)
	comps := PrepareComputations(xs[0], r, xs)
	c := RefractedColor(w, comps, 0)

	assert.True(t, ColorEquals(NewColor(0, 0, 0), c))
}

func TestRefractedColorUnderTotalInternalReflection(t *testing.T) {
	w := NewDefaultWorld()
	shape := w.Objects[0]
	shapeMaterial := shape.GetMaterial()
	shapeMaterial.Transparency = 1.0
	shapeMaterial.RefractiveIndex = 1.5
	shape.SetMaterial(shapeMaterial)
	r := NewRay(NewPoint(0, 0, 1/math.Sqrt(2)), NewVector(0, 1, 0))

	intersections := []Intersection{
		NewIntersection(-1/math.Sqrt(2), shape),
		NewIntersection(1/math.Sqrt(2), shape),
	}

	xs := NewIntersections(intersections...)
	comps := PrepareComputations(xs[1], r, xs)
	c := RefractedColor(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0, 0, 0), c))
}

func TestRefractedColorWithRefractedRay(t *testing.T) {
	w := NewDefaultWorld()

	A := w.Objects[0]
	materialA := A.GetMaterial()
	materialA.Ambient = 1.0
	materialA.SetPattern(NewTestPattern())
	A.SetMaterial(materialA)

	B := w.Objects[1]
	materialB := B.GetMaterial()
	materialB.Transparency = 1.0
	materialB.RefractiveIndex = 1.5
	B.SetMaterial(materialB)

	r := NewRay(NewPoint(0, 0, 0.1), NewVector(0, 1, 0))
	intersections := []Intersection{
		NewIntersection(-0.9899, A),
		NewIntersection(-0.4899, B),
		NewIntersection(0.4899, B),
		NewIntersection(0.9899, A),
	}

	xs := NewIntersections(intersections...)
	comps := PrepareComputations(xs[2], r, xs)
	c := RefractedColor(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0, 0.99888, 0.04725), c))
}

func TestShadeHitWithTransparentMaterial(t *testing.T) {
	w := NewDefaultWorld()
	floor := NewPlane()
	floor.SetTransform(Translate(0, -1, 0))
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5
	w.Objects = append(w.Objects, floor)

	ball := NewSphere()
	ball.SetTransform(Translate(0, -3.5, -0.5))
	ball.Material.Color = NewColor(1, 0, 0)
	ball.Material.Ambient = 0.5
	w.Objects = append(w.Objects, ball)

	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -1/math.Sqrt(2), 1/math.Sqrt(2)))
	xs := NewIntersections(NewIntersection(math.Sqrt(2), floor))
	comps := PrepareComputations(xs[0], r, xs)
	color := ShadeHit(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0.93642, 0.68642, 0.68642), color))
}

func TestSchlickApproximationUnderTotalInternalReflection(t *testing.T) {
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, 1/math.Sqrt(2)), NewVector(0, 1, 0))

	intersections := []Intersection{
		NewIntersection(-1/math.Sqrt(2), shape),
		NewIntersection(1/math.Sqrt(2), shape),
	}
	xs := NewIntersections(intersections...)

	comps := PrepareComputations(xs[1], r, xs)
	reflectance := Schlick(comps)

	assert.InDelta(t, 1.0, reflectance, float64EqualityThreshold)
}

func TestSchlickApproximationPerpendicularViewingAngle(t *testing.T) {
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))

	intersections := []Intersection{
		NewIntersection(-1, shape),
		NewIntersection(1, shape),
	}
	xs := NewIntersections(intersections...)

	comps := PrepareComputations(xs[1], r, xs)
	reflectance := Schlick(comps)

	assert.InDelta(t, 0.04, reflectance, float64EqualityThreshold)
}

func TestSchlickApproximationSmallAngleSecondIndexGreater(t *testing.T) {
	shape := NewGlassSphere()
	r := NewRay(NewPoint(0, 0.99, -2), NewVector(0, 0, 1))
	xs := NewIntersections(NewIntersection(1.8589, shape))
	comps := PrepareComputations(xs[0], r, xs)
	reflectance := Schlick(comps)

	assert.InDelta(t, 0.48873, reflectance, float64EqualityThreshold)
}

func TestShadeHitWithReflectiveTransparentMaterial(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -1/math.Sqrt(2), 1/math.Sqrt(2)))

	floor := NewPlane()
	floor.SetTransform(Translate(0, -1, 0))
	floor.Material.Reflective = 0.5
	floor.Material.Transparency = 0.5
	floor.Material.RefractiveIndex = 1.5
	w.Objects = append(w.Objects, floor)

	ball := NewSphere()
	ball.SetTransform(Translate(0, -3.5, -0.5))
	ball.Material.SetColor(NewColor(1, 0, 0))
	ball.Material.Ambient = 0.5
	w.Objects = append(w.Objects, ball)

	xs := NewIntersections(NewIntersection(math.Sqrt(2), floor))
	comps := PrepareComputations(xs[0], r, xs)
	color := ShadeHit(w, comps, RecursionDepth)

	assert.True(t, ColorEquals(NewColor(0.93391, 0.69643, 0.69243), color))
}

func TestIntersectionEncapsulatesUV(t *testing.T) {
	s := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))
	i := NewIntersectionUV(3.5, s, 0.2, 0.4)

	assert.InDelta(t, 0.2, i.U, float64EqualityThreshold)
	assert.InDelta(t, 0.4, i.V, float64EqualityThreshold)
}

func TestIntersectRayWithBoundingBoxAtOrigin(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
		result    bool
	}{
		{NewPoint(5, 0.5, 0), NewVector(-1, 0, 0), true},
		{NewPoint(-5, 0.5, 0), NewVector(1, 0, 0), true},
		{NewPoint(0.5, 5, 0), NewVector(0, -1, 0), true},
		{NewPoint(0.5, -5, 0), NewVector(0, 1, 0), true},
		{NewPoint(0.5, 0, 5), NewVector(0, 0, -1), true},
		{NewPoint(0.5, 0, -5), NewVector(0, 0, 1), true},
		{NewPoint(0, 0.5, 0), NewVector(0, 0, 1), true},
		{NewPoint(-2, 0, 0), NewVector(2, 4, 6), false},
		{NewPoint(0, -2, 0), NewVector(6, 2, 4), false},
		{NewPoint(0, 0, -2), NewVector(4, 6, 2), false},
		{NewPoint(2, 0, 2), NewVector(0, 0, -1), false},
		{NewPoint(0, 2, 2), NewVector(0, -1, 0), false},
		{NewPoint(2, 2, 0), NewVector(-1, 0, 0), false},
	}

	box := NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))

	for _, test := range testCases {
		direction := Normalize(test.direction)
		r := NewRay(test.origin, direction)

		assert.Equal(t, test.result, RayIntersectsBox(box, r))
	}
}

func TestIntersectRayWithNonCubicBoundingBox(t *testing.T) {
	testCases := []struct {
		origin    Tuple
		direction Tuple
		result    bool
	}{
		{NewPoint(15, 1, 2), NewVector(-1, 0, 0), true},
		{NewPoint(-5, -1, 4), NewVector(1, 0, 0), true},
		{NewPoint(7, 6, 5), NewVector(0, -1, 0), true},
		{NewPoint(9, -5, 6), NewVector(0, 1, 0), true},
		{NewPoint(8, 2, 12), NewVector(0, 0, -1), true},
		{NewPoint(6, 0, -5), NewVector(0, 0, 1), true},
		{NewPoint(8, 1, 3.5), NewVector(0, 0, 1), true},
		{NewPoint(9, -1, -8), NewVector(2, 4, 6), false},
		{NewPoint(8, 3, -4), NewVector(6, 2, 4), false},
		{NewPoint(9, -1, -2), NewVector(4, 6, 2), false},
		{NewPoint(4, 0, 9), NewVector(0, 0, -1), false},
		{NewPoint(8, 6, -1), NewVector(0, -1, 0), false},
		{NewPoint(12, 5, 4), NewVector(-1, 0, 0), false},
	}

	box := NewBoundingBox(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

	for _, test := range testCases {
		direction := Normalize(test.direction)
		r := NewRay(test.origin, direction)

		assert.Equal(t, test.result, RayIntersectsBox(box, r))
	}
}

func TestIntersectingRayAndGroupDoesNotTestChildrenIfBoxIsMisssed(t *testing.T) {
	child := NewTestShape()
	shape := NewGroup()
	shape.AddChild(child)

	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	Intersect(shape, r)
	assert.Equal(t, Ray{}, child.SavedRay)
}

func TestIntersectingRayAndGroupTestsChildrenIfBoxIsHit(t *testing.T) {
	child := NewTestShape()
	shape := NewGroup()
	shape.AddChild(child)

	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	Intersect(shape, r)
	assert.NotEqual(t, Ray{}, child.SavedRay)
}

func TestIntersectingRayAndCSGDoesNotTestChildrenIfBoxIsMissed(t *testing.T) {
	left := NewTestShape()
	right := NewTestShape()
	shape := NewCSG(CSGDifference, left, right)

	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))
	Intersect(shape, r)

	assert.Equal(t, Ray{}, left.SavedRay)
	assert.Equal(t, Ray{}, right.SavedRay)
}

func TestIntersectingRayAndCSGTestsChildrenIfBoxIsHit(t *testing.T) {
	left := NewTestShape()
	right := NewTestShape()
	shape := NewCSG(CSGDifference, left, right)

	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	Intersect(shape, r)

	assert.NotEqual(t, Ray{}, left.SavedRay)
	assert.NotEqual(t, Ray{}, right.SavedRay)
}
