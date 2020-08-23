package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointTranslation(t *testing.T) {
	transform := Translate(5, -3, 2)
	p := Point(-3, 4, 5)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 1.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 7.0, result.Z, float64EqualityThreshold)
}

func TestPointInverseTranslation(t *testing.T) {
	transform := Translate(5, -3, 2)
	invTransform := MatrixInverse(transform)
	p := Point(-3, 4, 5)
	result := MatrixTupleMultiply(invTransform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, -8.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 7.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 3.0, result.Z, float64EqualityThreshold)
}

func TestVectorNotAffectedByTranslation(t *testing.T) {
	transform := Translate(5, -3, 2)
	v := Vector(-3, 4, 5)
	result := MatrixTupleMultiply(transform, v)

	assert.True(t, result.IsVector())
	assert.InDelta(t, -3.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 4.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 5.0, result.Z, float64EqualityThreshold)
}

func TestPointScaling(t *testing.T) {
	transform := Scale(2, 3, 4)
	p := Point(-4, 6, 8)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, -8.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 18.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 32.0, result.Z, float64EqualityThreshold)
}

func TestVectorScaling(t *testing.T) {
	transform := Scale(2, 3, 4)
	v := Vector(-4, 6, 8)
	result := MatrixTupleMultiply(transform, v)

	assert.True(t, result.IsVector())
	assert.InDelta(t, -8.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 18.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 32.0, result.Z, float64EqualityThreshold)
}

func TestInverseScaling(t *testing.T) {
	transform := Scale(2, 3, 4)
	invTransform := MatrixInverse(transform)
	v := Vector(-4, 6, 8)
	result := MatrixTupleMultiply(invTransform, v)

	assert.True(t, result.IsVector())
	assert.InDelta(t, -2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 2.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 2.0, result.Z, float64EqualityThreshold)
}

func TestReflectionIsScalingByNegative(t *testing.T) {
	transform := Scale(-1, 1, 1)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, -2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 3.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 4.0, result.Z, float64EqualityThreshold)
}

func TestRotationAroundX(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarter := RotateX(math.Pi / 4)
	fullQuarter := RotateX(math.Pi / 2)

	halfQuarterResult := MatrixTupleMultiply(halfQuarter, p)
	assert.True(t, halfQuarterResult.IsPoint())
	assert.InDelta(t, 0.0, halfQuarterResult.X, float64EqualityThreshold)
	assert.InDelta(t, math.Sqrt(2)/2.0, halfQuarterResult.Y, float64EqualityThreshold)
	assert.InDelta(t, math.Sqrt(2)/2.0, halfQuarterResult.Z, float64EqualityThreshold)

	fullQuarterResult := MatrixTupleMultiply(fullQuarter, p)
	assert.True(t, fullQuarterResult.IsPoint())
	assert.InDelta(t, 0.0, fullQuarterResult.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, fullQuarterResult.Y, float64EqualityThreshold)
	assert.InDelta(t, 1.0, fullQuarterResult.Z, float64EqualityThreshold)
}

func TestRotationInverseRotatesInOppositeDirection(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarter := RotateX(math.Pi / 4)
	invHalfQuarter := MatrixInverse(halfQuarter)
	result := MatrixTupleMultiply(invHalfQuarter, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 0.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, math.Sqrt(2)/2.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, -math.Sqrt(2)/2.0, result.Z, float64EqualityThreshold)
}

func TestRotationAroundY(t *testing.T) {
	p := Point(0, 0, 1)

	halfQuarter := RotateY(math.Pi / 4)
	fullQuarter := RotateY(math.Pi / 2)

	halfQuarterResult := MatrixTupleMultiply(halfQuarter, p)
	assert.True(t, halfQuarterResult.IsPoint())
	assert.InDelta(t, math.Sqrt(2)/2.0, halfQuarterResult.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, halfQuarterResult.Y, float64EqualityThreshold)
	assert.InDelta(t, math.Sqrt(2)/2.0, halfQuarterResult.Z, float64EqualityThreshold)

	fullQuarterResult := MatrixTupleMultiply(fullQuarter, p)
	assert.True(t, fullQuarterResult.IsPoint())
	assert.InDelta(t, 1.0, fullQuarterResult.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, fullQuarterResult.Y, float64EqualityThreshold)
	assert.InDelta(t, 0.0, fullQuarterResult.Z, float64EqualityThreshold)
}

func TestRotationAroundZ(t *testing.T) {
	p := Point(0, 1, 0)

	halfQuarter := RotateZ(math.Pi / 4)
	fullQuarter := RotateZ(math.Pi / 2)

	halfQuarterResult := MatrixTupleMultiply(halfQuarter, p)
	assert.True(t, halfQuarterResult.IsPoint())
	assert.InDelta(t, -math.Sqrt(2)/2.0, halfQuarterResult.X, float64EqualityThreshold)
	assert.InDelta(t, math.Sqrt(2)/2.0, halfQuarterResult.Y, float64EqualityThreshold)
	assert.InDelta(t, 0, halfQuarterResult.Z, float64EqualityThreshold)

	fullQuarterResult := MatrixTupleMultiply(fullQuarter, p)
	assert.True(t, fullQuarterResult.IsPoint())
	assert.InDelta(t, -1.0, fullQuarterResult.X, float64EqualityThreshold)
	assert.InDelta(t, 0.0, fullQuarterResult.Y, float64EqualityThreshold)
	assert.InDelta(t, 0.0, fullQuarterResult.Z, float64EqualityThreshold)
}

func TestShearingXY(t *testing.T) {
	transform := Shear(1, 0, 0, 0, 0, 0)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 5.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 3.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 4.0, result.Z, float64EqualityThreshold)
}

func TestShearingXZ(t *testing.T) {
	transform := Shear(0, 1, 0, 0, 0, 0)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 6.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 3.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 4.0, result.Z, float64EqualityThreshold)
}

func TestShearingYX(t *testing.T) {
	transform := Shear(0, 0, 1, 0, 0, 0)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 5.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 4.0, result.Z, float64EqualityThreshold)
}

func TestShearingYZ(t *testing.T) {
	transform := Shear(0, 0, 0, 1, 0, 0)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 7.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 4.0, result.Z, float64EqualityThreshold)
}

func TestShearingZX(t *testing.T) {
	transform := Shear(0, 0, 0, 0, 1, 0)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 3.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 6.0, result.Z, float64EqualityThreshold)
}

func TestShearingZY(t *testing.T) {
	transform := Shear(0, 0, 0, 0, 0, 1)
	p := Point(2, 3, 4)
	result := MatrixTupleMultiply(transform, p)

	assert.True(t, result.IsPoint())
	assert.InDelta(t, 2.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 3.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 7.0, result.Z, float64EqualityThreshold)
}

func TestApplyingTransformationsInSequence(t *testing.T) {
	p := Point(1, 0, 1)
	matA := RotateX(math.Pi / 2)
	matB := Scale(5, 5, 5)
	matC := Translate(10, 5, 7)

	p2 := MatrixTupleMultiply(matA, p)
	assert.True(t, p2.Equals(Point(1, -1, 0)))

	p3 := MatrixTupleMultiply(matB, p2)
	assert.True(t, p3.Equals(Point(5, -5, 0)))

	p4 := MatrixTupleMultiply(matC, p3)
	assert.True(t, p4.Equals(Point(15, 0, 7)))
}

func TestChainingTransformationsInReverseOrder(t *testing.T) {
	p := Point(1, 0, 1)
	matA := RotateX(math.Pi / 2)
	matB := Scale(5, 5, 5)
	matC := Translate(10, 5, 7)
	transform := MatrixMultiply(MatrixMultiply(matC, matB), matA)

	result := MatrixTupleMultiply(transform, p)
	assert.True(t, result.Equals(Point(15, 0, 7)))
}
