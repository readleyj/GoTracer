package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalOnSphereAtAxisX(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, NewPoint(1, 0, 0), Intersection{})

	assert.True(t, n.IsVector())
	assert.InDelta(t, 1.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, n.Y, float64EqualityThreshold)
	assert.InDelta(t, 0.0, n.Z, float64EqualityThreshold)
}

func TestNormalOnSphereAtAxisY(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, NewPoint(0, 1, 0), Intersection{})

	assert.True(t, n.IsVector())
	assert.InDelta(t, 0.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 1.0, n.Y, float64EqualityThreshold)
	assert.InDelta(t, 0.0, n.Z, float64EqualityThreshold)
}

func TestNormalOnSphereAtAxisZ(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, NewPoint(0, 0, 1), Intersection{})

	assert.True(t, n.IsVector())
	assert.InDelta(t, 0.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, n.Y, float64EqualityThreshold)
	assert.InDelta(t, 1.0, n.Z, float64EqualityThreshold)
}

func TestNormalOnSphereAtNonaxicalPoint(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, NewPoint(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3)), Intersection{})

	assert.True(t, n.IsVector())
	assert.InDelta(t, 1/math.Sqrt(3), n.X, float64EqualityThreshold)
	assert.InDelta(t, 1/math.Sqrt(3), n.Y, float64EqualityThreshold)
	assert.InDelta(t, 1/math.Sqrt(3), n.Z, float64EqualityThreshold)
}

func TestNormalIsNormalizedVector(t *testing.T) {
	s := NewSphere()
	n := NormalAt(s, NewPoint(1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3)), Intersection{})
	normalized := Normalize(n)

	assert.True(t, TupleEquals(normalized, n))
}

func TestComputeNormalOnTranslatedSphere(t *testing.T) {
	s := NewSphere()
	s.SetTransform(Translate(0, 1, 0))
	n := NormalAt(s, NewPoint(0, 1.70711, -0.70711), Intersection{})

	assert.True(t, n.IsVector())
	assert.InDelta(t, 0.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 0.70711, n.Y, float64EqualityThreshold)
	assert.InDelta(t, -0.70711, n.Z, float64EqualityThreshold)
}

func TestComputeNormalOnTransformedSphere(t *testing.T) {
	s := NewSphere()
	m := MatrixMultiply(Scale(1, 0.5, 1), RotateZ(math.Pi/5))
	s.SetTransform(m)
	n := NormalAt(s, NewPoint(0, 1/math.Sqrt(2), -1/math.Sqrt(2)), Intersection{})

	assert.True(t, n.IsVector())
	assert.InDelta(t, 0.0, n.X, float64EqualityThreshold)
	assert.InDelta(t, 0.97014, n.Y, float64EqualityThreshold)
	assert.InDelta(t, -0.24254, n.Z, float64EqualityThreshold)
}

func TestReflectVector(t *testing.T) {
	v := NewVector(1, -1, 0)
	n := NewVector(0, 1, 0)
	r := Reflect(v, n)

	assert.True(t, r.IsVector())
	assert.InDelta(t, 1.0, r.X, float64EqualityThreshold)
	assert.InDelta(t, 1.0, r.Y, float64EqualityThreshold)
	assert.InDelta(t, 0.0, r.Z, float64EqualityThreshold)
}

func TestReflectVectorOffSlantedSurface(t *testing.T) {
	v := NewVector(0, -1, 0)
	n := NewVector(1/math.Sqrt(2), 1/math.Sqrt(2), 0)
	r := Reflect(v, n)

	assert.True(t, r.IsVector())
	assert.InDelta(t, 1.0, r.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, r.Y, float64EqualityThreshold)
	assert.InDelta(t, 0.0, r.Z, float64EqualityThreshold)
}
