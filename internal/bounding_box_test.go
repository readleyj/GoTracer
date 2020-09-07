package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEmptyBoundingBox(t *testing.T) {
	box := NewEmptyBoundingBox()

	assert.True(t, TupleEquals(NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)), box.Min))
	assert.True(t, TupleEquals(NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)), box.Max))
}

func TestCreateBoundingBoxWithVolume(t *testing.T) {
	box := NewBoundingBox(NewPoint(-1, -2, -3), NewPoint(3, 2, 1))

	assert.True(t, TupleEquals(NewPoint(-1, -2, -3), box.Min))
	assert.True(t, TupleEquals(NewPoint(3, 2, 1), box.Max))
}

func TestAddingPointsToEmptyBoundingBox(t *testing.T) {
	box := NewEmptyBoundingBox()
	p1 := NewPoint(-5, 2, 0)
	p2 := NewPoint(7, 0, -3)

	box.AddPoint(p1)
	box.AddPoint(p2)

	assert.True(t, TupleEquals(NewPoint(-5, 0, -3), box.Min))
	assert.True(t, TupleEquals(NewPoint(7, 2, 0), box.Max))
}

func TestSphereHasBoundingBox(t *testing.T) {
	shape := NewSphere()
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-1, -1, -1), box.Min))
	assert.True(t, TupleEquals(NewPoint(1, 1, 1), box.Max))
}

func TestPlaneHasBoundingBox(t *testing.T) {
	shape := NewPlane()
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(math.Inf(-1), 0, math.Inf(-1)), box.Min))
	assert.True(t, TupleEquals(NewPoint(math.Inf(1), 0, math.Inf(1)), box.Max))
}

func TestCubeHasBoundingBox(t *testing.T) {
	shape := NewCube()
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-1, -1, -1), box.Min))
	assert.True(t, TupleEquals(NewPoint(1, 1, 1), box.Max))
}

func TestUnboundedCylinderHasBoundingBox(t *testing.T) {
	shape := NewCylinder()
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-1, math.Inf(-1), -1), box.Min))
	assert.True(t, TupleEquals(NewPoint(1, math.Inf(1), 1), box.Max))
}

func TestBoundedCylinderHasBoundingBox(t *testing.T) {
	shape := NewCylinder()
	shape.Minimum = -5
	shape.Maximum = 3
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-1, -5, -1), box.Min))
	assert.True(t, TupleEquals(NewPoint(1, 3, 1), box.Max))
}

func TestUnboundedConeHasBoundingBox(t *testing.T) {
	shape := NewCone()
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)), box.Min))
	assert.True(t, TupleEquals(NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)), box.Max))
}

func TestBoundedConeHasBoundingBox(t *testing.T) {
	shape := NewCone()
	shape.Minimum = -5
	shape.Maximum = 3
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-5, -5, -5), box.Min))
	assert.True(t, TupleEquals(NewPoint(5, 3, 5), box.Max))
}

func TestTriangleHasBoundingBox(t *testing.T) {
	p1 := NewPoint(-3, 7, 2)
	p2 := NewPoint(6, 2, -4)
	p3 := NewPoint(2, -1, -1)
	shape := NewTriangle(p1, p2, p3)
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-3, -1, -4), box.Min))
	assert.True(t, TupleEquals(NewPoint(6, 7, 2), box.Max))
}

func TestTestShapeHasArbitraryBounds(t *testing.T) {
	shape := NewTestShape()
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-1, -1, -1), box.Min))
	assert.True(t, TupleEquals(NewPoint(1, 1, 1), box.Max))
}

func TestAddingBoundingBoxToAnother(t *testing.T) {
	box1 := NewBoundingBox(NewPoint(-5, -2, 0), NewPoint(7, 4, 4))
	box2 := NewBoundingBox(NewPoint(8, -7, -2), NewPoint(14, 2, 8))
	box1.AddBox(box2)

	assert.True(t, TupleEquals(NewPoint(-5, -7, -2), box1.Min))
	assert.True(t, TupleEquals(NewPoint(14, 4, 8), box1.Max))
}

func TestCheckIfBoxContainsGivenPoint(t *testing.T) {
	testCases := []struct {
		point  Tuple
		result bool
	}{
		{NewPoint(5, -2, 0), true},
		{NewPoint(11, 4, 7), true},
		{NewPoint(8, 1, 3), true},
		{NewPoint(3, 0, 3), false},
		{NewPoint(8, -4, 3), false},
		{NewPoint(8, 1, -1), false},
		{NewPoint(13, 1, 3), false},
		{NewPoint(8, 5, 3), false},
		{NewPoint(8, 1, 8), false},
	}

	box := NewBoundingBox(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

	for _, test := range testCases {
		assert.Equal(t, test.result, box.ContainsPoint(test.point))
	}
}

func TestCheckIfBoxContainsGivenBox(t *testing.T) {
	testCases := []struct {
		min    Tuple
		max    Tuple
		result bool
	}{
		{NewPoint(5, -2, 0), NewPoint(11, 4, 7), true},
		{NewPoint(6, -1, 1), NewPoint(10, 3, 6), true},
		{NewPoint(4, -3, -1), NewPoint(10, 3, 6), false},
		{NewPoint(6, -1, 1), NewPoint(12, 5, 8), false},
	}

	box := NewBoundingBox(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

	for _, test := range testCases {
		box2 := NewBoundingBox(test.min, test.max)
		assert.Equal(t, test.result, box.ContainsBox(box2))
	}
}

func TestTransformBoundingBox(t *testing.T) {
	box := NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
	matrix := MatrixMultiply(RotateX(math.Pi/4), RotateY(math.Pi/4))
	box2 := TransformBox(box, matrix)

	assert.True(t, TupleEquals(NewPoint(-1.4142, -1.7071, -1.7071), box2.Min))
	assert.True(t, TupleEquals(NewPoint(1.4142, 1.7071, 1.7071), box2.Max))
}

func TestQueryShapeBoundingBoxInParentSpace(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(MatrixMultiply(Translate(1, -3, 5), Scale(0.5, 2, 4)))
	box := ParentSpaceBoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(0.5, -5, 1), box.Min))
	assert.True(t, TupleEquals(NewPoint(1.5, -1, 9), box.Max))
}

func TestGroupHasBoundingBoxThatContainsChildren(t *testing.T) {
	s := NewSphere()
	s.SetTransform(MatrixMultiply(Translate(2, 5, -3), Scale(2, 2, 2)))

	c := NewCylinder()
	c.Minimum = -2
	c.Maximum = 2
	c.SetTransform(MatrixMultiply(Translate(-4, -1, 4), Scale(0.5, 1, 0.5)))

	shape := NewGroup()
	shape.AddChild(s)
	shape.AddChild(c)
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-4.5, -3, -5), box.Min))
	assert.True(t, TupleEquals(NewPoint(4, 7, 4.5), box.Max))
}

func TestCSGShapeHasBoundingBoxThatContainsItsChildren(t *testing.T) {
	left := NewSphere()

	right := NewSphere()
	right.SetTransform(Translate(2, 3, 4))

	shape := NewCSG(CSGDifference, left, right)
	box := BoundsOf(shape)

	assert.True(t, TupleEquals(NewPoint(-1, -1, -1), box.Min))
	assert.True(t, TupleEquals(NewPoint(3, 4, 5), box.Max))
}
