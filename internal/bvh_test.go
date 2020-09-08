package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitPerfectCube(t *testing.T) {
	box := NewBoundingBox(NewPoint(-1, -4, -5), NewPoint(9, 6, 5))
	left, right := SplitBounds(box)

	assert.True(t, TupleEquals(NewPoint(-1, -4, -5), left.Min))
	assert.True(t, TupleEquals(NewPoint(4, 6, 5), left.Max))
	assert.True(t, TupleEquals(NewPoint(4, -4, -5), right.Min))
	assert.True(t, TupleEquals(NewPoint(9, 6, 5), right.Max))
}

func TestSplitXWideBox(t *testing.T) {
	box := NewBoundingBox(NewPoint(-1, -2, -3), NewPoint(9, 5.5, 3))
	left, right := SplitBounds(box)

	assert.True(t, TupleEquals(NewPoint(-1, -2, -3), left.Min))
	assert.True(t, TupleEquals(NewPoint(4, 5.5, 3), left.Max))
	assert.True(t, TupleEquals(NewPoint(4, -2, -3), right.Min))
	assert.True(t, TupleEquals(NewPoint(9, 5.5, 3), right.Max))
}

func TestSplitYWideBox(t *testing.T) {
	box := NewBoundingBox(NewPoint(-1, -2, -3), NewPoint(5, 8, 3))
	left, right := SplitBounds(box)

	assert.True(t, TupleEquals(NewPoint(-1, -2, -3), left.Min))
	assert.True(t, TupleEquals(NewPoint(5, 3, 3), left.Max))
	assert.True(t, TupleEquals(NewPoint(-1, 3, -3), right.Min))
	assert.True(t, TupleEquals(NewPoint(5, 8, 3), right.Max))
}

func TestSplitZWideBox(t *testing.T) {
	box := NewBoundingBox(NewPoint(-1, -2, -3), NewPoint(5, 3, 7))
	left, right := SplitBounds(box)

	assert.True(t, TupleEquals(NewPoint(-1, -2, -3), left.Min))
	assert.True(t, TupleEquals(NewPoint(5, 3, 2), left.Max))
	assert.True(t, TupleEquals(NewPoint(-1, -2, 2), right.Min))
	assert.True(t, TupleEquals(NewPoint(5, 3, 7), right.Max))
}

func TestPartitionGroupChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(Translate(-2, 0, 0))

	s2 := NewSphere()
	s2.SetTransform(Translate(2, 0, 0))

	s3 := NewSphere()

	g := NewGroup()
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	left, right := PartitionChildren(g)

	assert.Equal(t, 1, len(g.Children))
	assert.True(t, ShapeEquals(s3, g.Children[0]))
	assert.True(t, ShapeEquals(s1, left.Children[0]))
	assert.True(t, ShapeEquals(s2, right.Children[0]))
}

func TestCreateSubgroupFromListOfChildren(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	g := NewGroup()

	MakeSubgroup(g, s1, s2)
	subgroup := g.Children[0].(*Group)

	assert.Equal(t, 1, len(g.Children))
	assert.Equal(t, 2, len(subgroup.Children))
	assert.True(t, ShapeEquals(s1, subgroup.Children[0]))
	assert.True(t, ShapeEquals(s2, subgroup.Children[1]))
}

func TestSubdividingPrimitiveDoesNothing(t *testing.T) {
	shape := NewSphere()
	Divide(shape, 1)
	assert.IsType(t, &Sphere{}, shape)
}

func TestSubdivingGroupPartitionsItsChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(Translate(-2, -2, 0))

	s2 := NewSphere()
	s2.SetTransform(Translate(-2, 2, 0))

	s3 := NewSphere()
	s3.SetTransform(Scale(4, 4, 4))

	g := NewGroup()
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	Divide(g, 1)

	assert.True(t, ShapeEquals(s3, g.Children[0]))
	subgroup := g.Children[1].(*Group)

	assert.Equal(t, 2, len(subgroup.Children))
	assert.True(t, ShapeEquals(s1, subgroup.Children[0].(*Group).Children[0]))
	assert.True(t, ShapeEquals(s2, subgroup.Children[1].(*Group).Children[0]))
}

func TestSubdividingGroupWithTooFewChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(Translate(-2, 0, 0))

	s2 := NewSphere()
	s2.SetTransform(Translate(2, 1, 0))

	s3 := NewSphere()
	s3.SetTransform(Translate(2, -1, 0))

	subgroup := NewGroup()
	subgroup.AddChild(s1)
	subgroup.AddChild(s2)
	subgroup.AddChild(s3)

	s4 := NewSphere()

	g := NewGroup()
	g.AddChild(subgroup)
	g.AddChild(s4)

	Divide(g, 3)

	assert.True(t, ShapeEquals(subgroup, g.Children[0]))
	assert.True(t, ShapeEquals(s4, g.Children[1]))

	assert.Equal(t, 1, len(subgroup.Children[0].(*Group).Children))
	assert.True(t, ShapeEquals(s1, subgroup.Children[0].(*Group).Children[0]))

	assert.Equal(t, 2, len(subgroup.Children[1].(*Group).Children))
	assert.True(t, ShapeEquals(s2, subgroup.Children[1].(*Group).Children[0]))
	assert.True(t, ShapeEquals(s3, subgroup.Children[1].(*Group).Children[1]))
}

func TestSubdividingCSGShapeSubdividesItsChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(Translate(-1.5, 0, 0))

	s2 := NewSphere()
	s2.SetTransform(Translate(1.5, 0, 0))

	left := NewGroup()
	left.AddChild(s1)
	left.AddChild(s2)

	s3 := NewSphere()
	s3.SetTransform(Translate(0, 0, -1.5))

	s4 := NewSphere()
	s4.SetTransform(Translate(0, 0, 1.5))

	right := NewGroup()
	right.AddChild(s3)
	right.AddChild(s4)

	shape := NewCSG(CSGDifference, left, right)
	Divide(shape, 1)

	assert.Equal(t, 1, len(left.Children[0].(*Group).Children))
	assert.True(t, ShapeEquals(s1, left.Children[0].(*Group).Children[0]))

	assert.Equal(t, 1, len(left.Children[1].(*Group).Children))
	assert.True(t, ShapeEquals(s2, left.Children[1].(*Group).Children[0]))

	assert.Equal(t, 1, len(right.Children[0].(*Group).Children))
	assert.True(t, ShapeEquals(s3, right.Children[0].(*Group).Children[0]))

	assert.Equal(t, 1, len(right.Children[1].(*Group).Children))
	assert.True(t, ShapeEquals(s4, right.Children[1].(*Group).Children[0]))
}
