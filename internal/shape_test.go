package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShapeDefaultTransform(t *testing.T) {
	s := NewTestShape()

	assert.True(t, MatrixEquals(Identity4, s.Transform))
}

func TestShapeAssigningTransform(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(Translate(2, 3, 4))

	assert.True(t, MatrixEquals(Translate(2, 3, 4), s.Transform))
}

func TestShapeDefaultMaterial(t *testing.T) {
	s := NewTestShape()
	m := s.Material

	assert.True(t, MaterialEquals(NewDefaultMaterial(), m))
}

func TestShapeAssigningMaterial(t *testing.T) {
	s := NewTestShape()
	m := NewDefaultMaterial()
	m.Ambient = 1
	s.SetMaterial(m)
	shapeMat := s.GetMaterial()

	assert.True(t, ColorEquals(m.Color, shapeMat.Color))
	assert.InDelta(t, m.Ambient, shapeMat.Ambient, float64EqualityThreshold)
	assert.InDelta(t, m.Diffuse, shapeMat.Diffuse, float64EqualityThreshold)
	assert.InDelta(t, m.Specular, shapeMat.Specular, float64EqualityThreshold)
	assert.InDelta(t, m.Shininess, shapeMat.Shininess, float64EqualityThreshold)
}

func TestIntersectScaledShapeWithRay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.SetTransform(Scale(2, 2, 2))
	Intersect(s, r)

	assert.True(t, TupleEquals(NewPoint(0, 0, -2.5), s.SavedRay.Origin))
	assert.True(t, TupleEquals(NewVector(0, 0, 0.5), s.SavedRay.Direction))
}

func TestIntersectTranslatedShapeWithRay(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()
	s.SetTransform(Translate(5, 0, 0))
	Intersect(s, r)

	assert.True(t, TupleEquals(NewPoint(-5, 0, -5), s.SavedRay.Origin))
	assert.True(t, TupleEquals(NewVector(0, 0, 1), s.SavedRay.Direction))
}

func TestComputeNormalOnTranslatedShape(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(Translate(0, 1, 0))
	n := NormalAt(s, NewPoint(0, 1.70711, -0.70711))

	assert.True(t, n.IsVector())
	assert.InDelta(t, 0.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 0.70711, n.Y, float64EqualityThreshold)
	assert.InDelta(t, -0.70711, n.Z, float64EqualityThreshold)
}

func TestComputeNormal(t *testing.T) {
	s := NewTestShape()
	m := MatrixMultiply(Scale(1, 0.5, 1), RotateZ(math.Pi/5))
	s.SetTransform(m)
	n := NormalAt(s, NewPoint(0, math.Pi/2, -math.Pi/2))

	assert.True(t, n.IsVector())
	assert.InDelta(t, 0.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 0.97014, n.Y, float64EqualityThreshold)
	assert.InDelta(t, -0.24254, n.Z, float64EqualityThreshold)
}

func TestShapeHasParentAttribute(t *testing.T) {
	s := NewTestShape()

	assert.Nil(t, s.Parent)
}
