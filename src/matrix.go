package internal

import (
	"github.com/google/go-cmp/cmp"
)

var Identity4 = MakeMatrix4([]float64{
	1, 0, 0, 0,
	0, 1, 0, 0,
	0, 0, 1, 0,
	0, 0, 0, 1,
})

type Matrix4 [16]float64
type Matrix3 [9]float64
type Matrix2 [4]float64

func MakeMatrix4(elems []float64) Matrix4 {
	var matrix [16]float64
	copy(matrix[:], elems)

	return matrix
}

func Matrix4Equals(mat1 Matrix4, mat2 Matrix4) bool {
	return cmp.Equal(mat1, mat2, opt)
}

func Matrix4Multiply(mat1 Matrix4, mat2 Matrix4) Matrix4 {
	mat3 := MakeMatrix4(make([]float64, 16))

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			mat3.Set(multiplyRowCol(mat1, mat2, row, col), row, col)
		}
	}

	return mat3
}

func Matrix4TupleMultiply(mat Matrix4, t Tuple) Tuple {
	results := make([]float64, 4)

	for row := 0; row < 4; row++ {
		a0 := mat.Get(row, 0) * t.X
		a1 := mat.Get(row, 1) * t.Y
		a2 := mat.Get(row, 2) * t.Z
		a3 := mat.Get(row, 3) * t.W

		results[row] = a0 + a1 + a2 + a3
	}

	return Tuple{results[0], results[1], results[2], results[3]}
}

func Matrix4Transpose(mat Matrix4) Matrix4 {
	transposed := MakeMatrix4(make([]float64, 16))

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			transposed.Set(mat.Get(col, row), row, col)
		}
	}

	return transposed
}

func Matrix4Submatrix(mat Matrix4, row, col int) Matrix3 {
	submatrix := MakeMatrix3(make([]float64, 9))
	currIndex := 0

	for i := 0; i < 4; i++ {
		if i == row {
			continue
		}
		for j := 0; j < 4; j++ {
			if j == col {
				continue
			}
			submatrix[currIndex] = mat.Get(i, j)
			currIndex++
		}
	}

	return submatrix
}

func Matrix4Minor(mat Matrix4, row, col int) float64 {
	submatrix := Matrix4Submatrix(mat, row, col)
	return Matrix3Determinant(submatrix)
}

func Matrix4Cofactor(mat Matrix4, row, col int) float64 {
	cofactor := Matrix4Minor(mat, row, col)

	if (row+col)%2 != 0 {
		cofactor *= -1
	}

	return cofactor
}

func Matrix4Determinant(mat Matrix4) float64 {
	var determinant float64 = 0.0

	for col := 0; col < 4; col++ {
		determinant += (mat.Get(0, col) * Matrix4Cofactor(mat, 0, col))
	}

	return determinant
}

func Matrix4IsInvertible(mat Matrix4) bool {
	determinant := Matrix4Determinant(mat)

	return !(determinant == 0)
}

func Matrix4Inverse(mat Matrix4) Matrix4 {
	if !Matrix4IsInvertible(mat) {
		panic("The given matrix is not invertible")
	}

	inverse := MakeMatrix4(make([]float64, 16))
	determinant := Matrix4Determinant(mat)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			cofactor := Matrix4Cofactor(mat, row, col)
			inverse.Set(cofactor/determinant, col, row)
		}
	}

	return inverse
}

func (mat Matrix4) Get(row, col int) float64 {
	return mat[row*4+col]
}

func (mat *Matrix4) Set(val float64, row, col int) {
	mat[row*4+col] = val
}

func (mat Matrix4) GetByIndex(index int) float64 {
	return mat[index]
}

func MakeMatrix3(elems []float64) Matrix3 {
	var matrix [9]float64
	copy(matrix[:], elems)

	return matrix
}

func Matrix3Submatrix(mat Matrix3, row, col int) Matrix2 {
	submatrix := MakeMatrix2(make([]float64, 4))
	currIndex := 0

	for i := 0; i < 3; i++ {
		if i == row {
			continue
		}
		for j := 0; j < 3; j++ {
			if j == col {
				continue
			}

			submatrix[currIndex] = mat.Get(i, j)
			currIndex++
		}
	}

	return submatrix
}

func Matrix3Minor(mat Matrix3, row, col int) float64 {
	submatrix := Matrix3Submatrix(mat, row, col)
	return Matrix2Determinant(submatrix)
}

func Matrix3Cofactor(mat Matrix3, row, col int) float64 {
	cofactor := Matrix3Minor(mat, row, col)

	if (row+col)%2 != 0 {
		cofactor *= -1
	}

	return cofactor
}

func Matrix3Determinant(mat Matrix3) float64 {
	var determinant float64 = 0.0

	for col := 0; col < 3; col++ {
		determinant += (mat.Get(0, col) * Matrix3Cofactor(mat, 0, col))
	}

	return determinant
}

func Matrix3Equals(mat1 Matrix3, mat2 Matrix3) {
	cmp.Equal(mat1, mat2, opt)
}

func (mat Matrix3) Get(row, col int) float64 {
	return mat[row*3+col]
}

func (mat *Matrix3) Set(val float64, row, col int) {
	mat[row*3+col] = val
}

func (mat Matrix3) GetByIndex(index int) float64 {
	return mat[index]
}

func MakeMatrix2(elems []float64) Matrix2 {
	var matrix [4]float64
	copy(matrix[:], elems)

	return matrix
}

func Matrix2Equals(mat1 Matrix2, mat2 Matrix2) {
	cmp.Equal(mat1, mat2, opt)
}

func Matrix2Determinant(mat Matrix2) float64 {
	return mat.Get(0, 0)*mat.Get(1, 1) - mat.Get(0, 1)*mat.Get(1, 0)
}

func (mat Matrix2) Get(row, col int) float64 {
	return mat[row*2+col]
}

func (mat *Matrix2) Set(val float64, row, col int) {
	mat[row*2+col] = val
}

func (mat Matrix2) GetByIndex(index int) float64 {
	return mat[index]
}

func (mat Matrix2) Equals(mat1 Matrix2) bool {
	return cmp.Equal(mat, mat1, opt)
}

func multiplyRowCol(mat1, mat2 Matrix4, row, col int) float64 {
	a0 := mat1.Get(row, 0) * mat2.Get(0, col)
	a1 := mat1.Get(row, 1) * mat2.Get(1, col)
	a2 := mat1.Get(row, 2) * mat2.Get(2, col)
	a3 := mat1.Get(row, 3) * mat2.Get(3, col)

	return a0 + a1 + a2 + a3
}
