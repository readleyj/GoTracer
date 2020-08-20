package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTupleIsVector(t *testing.T) {
	v := Tuple{4.3, -4.2, 3.1, 0.0}

	assert.True(t, v.IsVector())
	assert.False(t, v.IsPoint())

	assert.InDelta(t, 4.3, v.X, float64EqualityThreshold)
	assert.InDelta(t, -4.2, v.Y, float64EqualityThreshold)
	assert.InDelta(t, 3.1, v.Z, float64EqualityThreshold)
}

func TestTupleIsPoint(t *testing.T) {
	p := Tuple{4.3, -4.2, 3.1, 1.0}

	assert.True(t, p.IsPoint())
	assert.False(t, p.IsVector())

	assert.InDelta(t, 4.3, p.X, float64EqualityThreshold)
	assert.InDelta(t, -4.2, p.Y, float64EqualityThreshold)
	assert.InDelta(t, 3.1, p.Z, float64EqualityThreshold)
}

func TestVectorIsTuple(t *testing.T) {
	v := Vector(4, -4, 3)

	assert.InDelta(t, 4, v.X, float64EqualityThreshold)
	assert.InDelta(t, -4, v.Y, float64EqualityThreshold)
	assert.InDelta(t, 3, v.Z, float64EqualityThreshold)
	assert.InDelta(t, 0.0, v.W, float64EqualityThreshold)
}

func TestPointIsTuple(t *testing.T) {
	p := Point(4, -4, 3)

	assert.InDelta(t, 4.0, p.X, float64EqualityThreshold)
	assert.InDelta(t, -4.0, p.Y, float64EqualityThreshold)
	assert.InDelta(t, 3.0, p.Z, float64EqualityThreshold)
	assert.InDelta(t, 1.0, p.W, float64EqualityThreshold)
}

func TestTupleAdd(t *testing.T) {
	t1 := Tuple{3, -2, 5, 1}
	t2 := Tuple{-2, 3, 1, 0}
	t3 := Add(t1, t2)

	assert.InDelta(t, 1.0, t3.X, float64EqualityThreshold)
	assert.InDelta(t, 1.0, t3.Y, float64EqualityThreshold)
	assert.InDelta(t, 6.0, t3.Z, float64EqualityThreshold)
	assert.InDelta(t, 1.0, t3.W, float64EqualityThreshold)
}

func TestTupleSub(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	p3 := Sub(p1, p2)

	assert.InDelta(t, -2.0, p3.X, float64EqualityThreshold)
	assert.InDelta(t, -4.0, p3.Y, float64EqualityThreshold)
	assert.InDelta(t, -6.0, p3.Z, float64EqualityThreshold)
	assert.InDelta(t, 0.0, p3.W, float64EqualityThreshold)
}

func TestVectorSub(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	v3 := Sub(v1, v2)

	assert.InDelta(t, -2.0, v3.X, float64EqualityThreshold)
	assert.InDelta(t, -4.0, v3.Y, float64EqualityThreshold)
	assert.InDelta(t, -6.0, v3.Z, float64EqualityThreshold)
	assert.InDelta(t, 0.0, v3.W, float64EqualityThreshold)
}

func TestSubVectorFromZero(t *testing.T) {
	v1 := Vector(0, 0, 0)
	v2 := Vector(1, -2, 3)
	v3 := Sub(v1, v2)

	assert.InDelta(t, -1.0, v3.X, float64EqualityThreshold)
	assert.InDelta(t, 2.0, v3.Y, float64EqualityThreshold)
	assert.InDelta(t, -3.0, v3.Z, float64EqualityThreshold)
	assert.InDelta(t, 0.0, v3.W, float64EqualityThreshold)
}

func TestTupleNegate(t *testing.T) {
	v1 := Tuple{1, -2, 3, -4}
	v2 := Negate(v1)

	assert.InDelta(t, -1.0, v2.X, float64EqualityThreshold)
	assert.InDelta(t, 2.0, v2.Y, float64EqualityThreshold)
	assert.InDelta(t, -3.0, v2.Z, float64EqualityThreshold)
	assert.InDelta(t, 4.0, v2.W, float64EqualityThreshold)
}

func TestMultiplyByScalar(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	t2 := MultiplyByScalar(t1, 3.5)

	assert.InDelta(t, 3.5, t2.X, float64EqualityThreshold)
	assert.InDelta(t, -7.0, t2.Y, float64EqualityThreshold)
	assert.InDelta(t, 10.5, t2.Z, float64EqualityThreshold)
	assert.InDelta(t, -14.0, t2.W, float64EqualityThreshold)
}

func TestMultiplyByFraction(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	t2 := MultiplyByScalar(t1, 0.5)

	assert.InDelta(t, 0.5, t2.X, float64EqualityThreshold)
	assert.InDelta(t, -1.0, t2.Y, float64EqualityThreshold)
	assert.InDelta(t, 1.5, t2.Z, float64EqualityThreshold)
	assert.InDelta(t, -2.0, t2.W, float64EqualityThreshold)
}

func TestDivideByScalar(t *testing.T) {
	t1 := Tuple{1, -2, 3, -4}
	t2 := DivideByScalar(t1, 2)

	assert.InDelta(t, 0.5, t2.X, float64EqualityThreshold)
	assert.InDelta(t, -1.0, t2.Y, float64EqualityThreshold)
	assert.InDelta(t, 1.5, t2.Z, float64EqualityThreshold)
	assert.InDelta(t, -2.0, t2.W, float64EqualityThreshold)
}

func TestVectorMagnitude(t *testing.T) {
	testCases := []struct {
		tuple  Tuple
		target float64
	}{
		{Vector(1, 0, 0), 1.0},
		{Vector(0, 1, 0), 1.0},
		{Vector(0, 0, 1), 1.0},
		{Vector(1, 2, 3), math.Sqrt(14)},
		{Vector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, test := range testCases {
		assert.InDelta(t, test.target, Magnitude(test.tuple), float64EqualityThreshold)
	}
}

func TestNormalizedVectorMagnitude(t *testing.T) {
	testCases := []struct {
		input  Tuple
		target Tuple
	}{
		{Vector(4, 0, 0), Vector(1, 0, 0)},
		{Vector(1, 2, 3), Vector(0.26726, 0.53452, 0.80178)},
	}

	for _, test := range testCases {
		normalized, target := Normalize(test.input), test.target

		assert.InDelta(t, target.X, normalized.X, float64EqualityThreshold)
		assert.InDelta(t, target.Y, normalized.Y, float64EqualityThreshold)
		assert.InDelta(t, target.Z, normalized.Z, float64EqualityThreshold)
		assert.InDelta(t, target.W, normalized.W, float64EqualityThreshold)
	}
}

func TestDotProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)
	result := Dot(v1, v2)
	assert.InDelta(t, 20.0, result, float64EqualityThreshold)
}

func TestCrossProduct(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	targetV12 := Vector(-1, 2, -1)
	targetV21 := Vector(1, -2, 1)

	assert.True(t, TupleEquals(targetV12, Cross(v1, v2)))
	assert.True(t, TupleEquals(targetV21, Cross(v2, v1)))
}
