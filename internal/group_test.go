package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	g := NewGroup()

	assert.True(t, MatrixEquals(Identity4, g.Transform))
}

func TestAddChildToGroup(t *testing.T) {
	g := NewGroup()
	s := NewTestShape()
	g.AddChild(s)

	_, includes := Includes(g.Children, s)

	assert.NotEqual(t, 0, len(g.Children))
	assert.True(t, includes)
	assert.True(t, ShapeEquals(s.Parent, g))
}

func TestIntersectRayWithEmptyGroup(t *testing.T) {
	g := NewGroup()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	xs := g.LocalIntersect(r)

	assert.Equal(t, 0, len(xs))
}

func TestIntersectRayWithNonemptyGroup(t *testing.T) {
	g := NewGroup()
	s1 := NewSphere()

	s2 := NewSphere()
	s2.SetTransform(Translate(0, 0, -3))

	s3 := NewSphere()
	s3.SetTransform(Translate(5, 0, 0))

	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := g.LocalIntersect(r)

	assert.Equal(t, 4, len(xs))
	assert.True(t, ShapeEquals(s2, xs[0].Object))
	assert.True(t, ShapeEquals(s2, xs[1].Object))
	assert.True(t, ShapeEquals(s1, xs[2].Object))
	assert.True(t, ShapeEquals(s1, xs[3].Object))
}

func TestIntersectTransformedGroup(t *testing.T) {
	g := NewGroup()
	g.SetTransform(Scale(2, 2, 2))

	s := NewSphere()
	s.SetTransform(Translate(5, 0, 0))
	g.AddChild(s)

	r := NewRay(NewPoint(10, 0, -10), NewVector(0, 0, 1))
	xs := Intersect(g, r)

	assert.Equal(t, 2, len(xs))
}

func TestConvertPointFromWorldToObjectSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotateY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(Scale(2, 2, 2))
	g1.AddChild(g2)

	s := NewSphere()
	s.SetTransform(Translate(5, 0, 0))
	g2.AddChild(s)

	p := WorldToObject(s, NewPoint(-2, 0, -10))

	assert.True(t, TupleEquals(NewPoint(0, 0, -1), p))
}

func TestConvertNormalFromObjectToWorldSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotateY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(Scale(1, 2, 3))
	g1.AddChild(g2)

	s := NewSphere()
	s.SetTransform(Translate(5, 0, 0))
	g2.AddChild(s)

	n := NormalToWorld(s, NewVector(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3)))

	assert.True(t, TupleEquals(NewVector(0.2857, 0.4286, -0.8571), n))
}

func TestNormalOnChildObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotateY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(Scale(1, 2, 3))
	g1.AddChild(g2)

	s := NewSphere()
	s.SetTransform(Translate(5, 0, 0))
	g2.AddChild(s)

	n := NormalAt(s, NewPoint(1.7321, 1.1547, -5.5774))

	assert.True(t, TupleEquals(NewVector(0.2857, 0.4286, -0.8571), n))
}
