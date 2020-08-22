package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMatrix4(t *testing.T) {
	elements := []float64{
		1, 2, 3, 4,
		5.5, 6.5, 7.5, 8.5,
		9, 10, 11, 12,
		13.5, 14.5, 15.5, 16.5,
	}

	mat := MakeMatrix4(elements)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			assert.InDelta(t, elements[i*4+j], mat.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestCreateMatrix3(t *testing.T) {
	elements := []float64{
		-3, 5, 0,
		1, -2, -7,
		0, 1, 1,
	}

	mat := MakeMatrix3(elements)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.InDelta(t, elements[i*3+j], mat.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestCreateMatrix2(t *testing.T) {
	elements := []float64{
		-3, 5,
		1, -2,
	}

	mat := MakeMatrix2(elements)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			assert.InDelta(t, elements[i*2+j], mat.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestMatrix4EqualityIdentical(t *testing.T) {
	elems1 := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}

	elems2 := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}

	mat1 := MakeMatrix4(elems1)
	mat2 := MakeMatrix4(elems2)

	assert.True(t, Matrix4Equals(mat1, mat2))
}

func TestMatrix4EqualityDifferent(t *testing.T) {
	elems1 := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}

	elems2 := []float64{
		2, 3, 4, 5,
		6, 7, 8, 9,
		8, 7, 6, 5,
		4, 3, 2, 1,
	}

	mat1 := MakeMatrix4(elems1)
	mat2 := MakeMatrix4(elems2)

	assert.False(t, Matrix4Equals(mat1, mat2))
}

func TestMatrixMultiply4x4(t *testing.T) {
	elems1 := []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 8, 7, 6,
		5, 4, 3, 2,
	}

	elems2 := []float64{
		-2, 1, 2, 3,
		3, 2, 1, -1,
		4, 3, 6, 5,
		1, 2, 7, 8,
	}

	elems3 := []float64{
		20, 22, 50, 48,
		44, 54, 114, 108,
		40, 58, 110, 102,
		16, 26, 46, 42,
	}

	mat1 := MakeMatrix4(elems1)
	mat2 := MakeMatrix4(elems2)
	mat3 := Matrix4Multiply(mat1, mat2)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			assert.InDelta(t, elems3[i*4+j], mat3.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestMatrix4TupleMultiply(t *testing.T) {
	elems := []float64{
		1, 2, 3, 4,
		2, 4, 4, 2,
		8, 6, 4, 1,
		0, 0, 0, 1,
	}

	mat := MakeMatrix4(elems)
	tuple := Tuple{1, 2, 3, 1}
	result := Matrix4TupleMultiply(mat, tuple)

	assert.InDelta(t, 18.0, result.X, float64EqualityThreshold)
	assert.InDelta(t, 24.0, result.Y, float64EqualityThreshold)
	assert.InDelta(t, 33.0, result.Z, float64EqualityThreshold)
	assert.InDelta(t, 1.0, result.W, float64EqualityThreshold)
}

func TestMatrix4IdentityMultiply(t *testing.T) {
	elems := []float64{
		0, 1, 2, 4,
		1, 2, 4, 8,
		2, 4, 8, 16,
		4, 8, 16, 32,
	}

	mat := MakeMatrix4(elems)
	identity := Identity4

	matByIdentity := Matrix4Multiply(mat, identity)
	identityByMat := Matrix4Multiply(identity, mat)

	assert.True(t, Matrix4Equals(mat, matByIdentity))
	assert.True(t, Matrix4Equals(mat, identityByMat))
}

func TestMatrix4Transpose(t *testing.T) {
	elems := []float64{
		0, 9, 3, 0,
		9, 8, 0, 8,
		1, 8, 5, 3,
		0, 0, 5, 8,
	}

	transposedElems := []float64{
		0, 9, 1, 0,
		9, 8, 8, 0,
		3, 0, 5, 5,
		0, 8, 3, 8,
	}

	mat1 := MakeMatrix4(elems)
	mat2 := Matrix4Transpose(mat1)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			assert.InDelta(t, transposedElems[i*4+j], mat2.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestIdentity4Transpose(t *testing.T) {
	transposedIdentity := Matrix4Transpose(Identity4)

	assert.True(t, Matrix4Equals(Identity4, transposedIdentity))
}

func TestMatrix2Determinant(t *testing.T) {
	elems := []float64{
		1, 5,
		-3, 2,
	}

	mat := MakeMatrix2(elems)
	determinant := Matrix2Determinant(mat)

	assert.InDelta(t, 17.0, determinant, float64EqualityThreshold)
}

func TestMatrix4Submatrix(t *testing.T) {
	elems := []float64{
		-6, 1, 1, 6,
		-8, 5, 8, 6,
		-1, 0, 8, 2,
		-7, 1, -1, 1,
	}

	targetElems := []float64{
		-6, 1, 6,
		-8, 8, 6,
		-7, -1, 1,
	}

	mat := MakeMatrix4(elems)
	submatrix := Matrix4Submatrix(mat, 2, 1)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			assert.InDelta(t, targetElems[i*3+j], submatrix.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestMatrix3Submatrix(t *testing.T) {
	elems := []float64{
		1, 5, 0,
		-3, 2, 7,
		0, 6, -3,
	}

	targetElems := []float64{
		-3, 2,
		0, 6,
	}

	mat := MakeMatrix3(elems)
	submatrix := Matrix3Submatrix(mat, 0, 2)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			assert.InDelta(t, targetElems[i*2+j], submatrix.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestMatrix3Minor(t *testing.T) {
	elems := []float64{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	}

	mat := MakeMatrix3(elems)
	matB := Matrix3Submatrix(mat, 1, 0)

	assert.InDelta(t, 25.0, Matrix2Determinant(matB), float64EqualityThreshold)
	assert.InDelta(t, 25.0, Matrix3Minor(mat, 1, 0), float64EqualityThreshold)
}

func TestMatrix3Cofactor(t *testing.T) {
	elems := []float64{
		3, 5, 0,
		2, -1, -7,
		6, -1, 5,
	}

	mat := MakeMatrix3(elems)

	minor1, cofactor1 := Matrix3Minor(mat, 0, 0), Matrix3Cofactor(mat, 0, 0)
	assert.InDelta(t, -12.0, minor1, float64EqualityThreshold)
	assert.InDelta(t, -12.0, cofactor1, float64EqualityThreshold)

	minor2, cofactor2 := Matrix3Minor(mat, 1, 0), Matrix3Cofactor(mat, 1, 0)
	assert.InDelta(t, 25.0, minor2, float64EqualityThreshold)
	assert.InDelta(t, -25.0, cofactor2, float64EqualityThreshold)
}

func TestMatrix3Determinant(t *testing.T) {
	elems := []float64{
		1, 2, 6,
		-5, 8, -4,
		2, 6, 4,
	}

	mat := MakeMatrix3(elems)

	cofactor1 := Matrix3Cofactor(mat, 0, 0)
	cofactor2 := Matrix3Cofactor(mat, 0, 1)
	cofactor3 := Matrix3Cofactor(mat, 0, 2)
	determinant := Matrix3Determinant(mat)

	assert.InDelta(t, 56.0, cofactor1, float64EqualityThreshold)
	assert.InDelta(t, 12.0, cofactor2, float64EqualityThreshold)
	assert.InDelta(t, -46.0, cofactor3, float64EqualityThreshold)
	assert.InDelta(t, -196.0, determinant, float64EqualityThreshold)
}

func TestMatrix4Determinant(t *testing.T) {
	elems := []float64{
		-2, -8, 3, 5,
		-3, 1, 7, 3,
		1, 2, -9, 6,
		-6, 7, 7, -9,
	}

	mat := MakeMatrix4(elems)

	cofactor1 := Matrix4Cofactor(mat, 0, 0)
	cofactor2 := Matrix4Cofactor(mat, 0, 1)
	cofactor3 := Matrix4Cofactor(mat, 0, 2)
	cofactor4 := Matrix4Cofactor(mat, 0, 3)
	determinant := Matrix4Determinant(mat)

	assert.InDelta(t, 690.0, cofactor1, float64EqualityThreshold)
	assert.InDelta(t, 447.0, cofactor2, float64EqualityThreshold)
	assert.InDelta(t, 210.0, cofactor3, float64EqualityThreshold)
	assert.InDelta(t, 51.0, cofactor4, float64EqualityThreshold)
	assert.InDelta(t, -4071.0, determinant, float64EqualityThreshold)
}

func TestInvertabilityForInvertable(t *testing.T) {
	elems := []float64{
		6, 4, 4, 4,
		5, 5, 7, 6,
		4, -9, 3, -7,
		9, 1, 7, -6,
	}

	mat := MakeMatrix4(elems)

	determinant := Matrix4Determinant(mat)
	assert.InDelta(t, -2120.0, determinant, float64EqualityThreshold)
	assert.True(t, Matrix4IsInvertible(mat))
}

func TestInvertabilityForNoninvertable(t *testing.T) {
	elems := []float64{
		-4, 2, -2, -3,
		9, 6, 2, 6,
		0, -5, 1, -5,
		0, 0, 0, 0,
	}

	mat := MakeMatrix4(elems)

	determinant := Matrix4Determinant(mat)
	assert.InDelta(t, 0.0, determinant, float64EqualityThreshold)
	assert.False(t, Matrix4IsInvertible(mat))
}

func TestMatrix4Inverse(t *testing.T) {
	elems := []float64{
		-5, 2, 6, -8,
		1, -5, 1, 8,
		7, 7, -6, -7,
		1, -3, 7, 4,
	}

	targetElems := []float64{
		0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639,
	}

	mat := MakeMatrix4(elems)
	inverse := Matrix4Inverse(mat)

	determinant := Matrix4Determinant(mat)
	cofactor1 := Matrix4Cofactor(mat, 2, 3)
	cofactor2 := Matrix4Cofactor(mat, 3, 2)

	assert.InDelta(t, 532.0, determinant, float64EqualityThreshold)
	assert.InDelta(t, -160.0, cofactor1, float64EqualityThreshold)
	assert.InDelta(t, -160.0/532.0, inverse.Get(3, 2), float64EqualityThreshold)
	assert.InDelta(t, 105.0, cofactor2, float64EqualityThreshold)
	assert.InDelta(t, 105.0/532.0, inverse.Get(2, 3), float64EqualityThreshold)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			assert.InDelta(t, targetElems[i*4+j], inverse.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestMatrix4InverseAdditional(t *testing.T) {
	elems := []float64{
		8, -5, 9, 2,
		7, 5, 6, 1,
		-6, 0, 9, 6,
		-3, 0, -9, -4,
	}

	targetElems := []float64{
		-0.15385, -0.15385, -0.28205, -0.53846,
		-0.07692, 0.12308, 0.02564, 0.03077,
		0.35897, 0.35897, 0.43590, 0.92308,
		-0.69231, -0.69231, -0.76923, -1.92308,
	}

	mat := MakeMatrix4(elems)
	inverse := Matrix4Inverse(mat)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			assert.InDelta(t, targetElems[i*4+j], inverse.Get(i, j), float64EqualityThreshold)
		}
	}

	elems = []float64{
		9, 3, 0, 9,
		-5, -2, -6, -3,
		-4, 9, 6, 4,
		-7, 6, 6, 2,
	}

	targetElems = []float64{
		-0.04074, -0.07778, 0.14444, -0.22222,
		-0.07778, 0.03333, 0.36667, -0.33333,
		-0.02901, -0.14630, -0.10926, 0.12963,
		0.17778, 0.06667, -0.26667, 0.33333,
	}

	mat = MakeMatrix4(elems)
	inverse = Matrix4Inverse(mat)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			assert.InDelta(t, targetElems[i*4+j], inverse.Get(i, j), float64EqualityThreshold)
		}
	}
}

func TestInverseReversesMultiplication(t *testing.T) {
	elemsA := []float64{
		3, -9, 7, 3,
		3, -8, 2, -9,
		-4, 4, 4, 1,
		-6, 5, -1, 1,
	}

	elemsB := []float64{
		8, 2, 2, 2,
		3, -1, 7, 0,
		7, 0, 5, 4,
		6, -2, 0, 5,
	}

	matA := MakeMatrix4(elemsA)
	matB := MakeMatrix4(elemsB)
	matC := Matrix4Multiply(matA, matB)
	reversed := Matrix4Multiply(matC, Matrix4Inverse(matB))

	assert.True(t, Matrix4Equals(matA, reversed))
}
