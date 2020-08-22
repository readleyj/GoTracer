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

type Matrix struct {
	Rows, Cols int
	elems      []float64
}

func MakeMatrix(elems []float64, numRows, numCols int) Matrix {
	if len(elems) != numRows*numCols {
		panic("The number of values and matrix dimensions do not match")
	}

	mat := Matrix{numRows, numCols, make([]float64, numRows*numCols)}
	copy(mat.elems[:], elems)
	return mat
}

func MatrixEquals(mat1, mat2 Matrix) bool {
	return cmp.Equal(mat1.elems, mat2.elems, opt)
}

func MakeMatrix4(elems []float64) Matrix {
	return MakeMatrix(elems, 4, 4)
}

func MakeMatrix3(elems []float64) Matrix {
	return MakeMatrix(elems, 3, 3)

}

func MakeMatrix2(elems []float64) Matrix {
	return MakeMatrix(elems, 2, 2)
}

func MatrixMultiply(mat1, mat2 Matrix) Matrix {
	mat3 := MakeMatrix4(make([]float64, mat1.Rows*mat2.Cols))

	for row := 0; row < mat1.Rows; row++ {
		for col := 0; col < mat2.Cols; col++ {
			mat3.Set(multiplyRowCol(mat1, mat2, row, col), row, col)
		}
	}

	return mat3
}

func MatrixTupleMultiply(mat Matrix, t Tuple) Tuple {
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

func MatrixTranspose(mat Matrix) Matrix {
	transposed := MakeMatrix4(make([]float64, mat.Rows*mat.Cols))

	for row := 0; row < mat.Rows; row++ {
		for col := 0; col < mat.Cols; col++ {
			transposed.Set(mat.Get(col, row), row, col)
		}
	}

	return transposed
}

func MatrixSubmatrix(mat Matrix, row, col int) Matrix {
	submatrix := MakeMatrix(make([]float64, (mat.Rows-1)*(mat.Cols-1)), mat.Rows-1, mat.Cols-1)
	currIndex := 0

	for i := 0; i < mat.Rows; i++ {
		if i == row {
			continue
		}
		for j := 0; j < mat.Cols; j++ {
			if j == col {
				continue
			}
			submatrix.elems[currIndex] = mat.Get(i, j)
			currIndex++
		}
	}

	return submatrix
}

func MatrixMinor(mat Matrix, row, col int) float64 {
	submatrix := MatrixSubmatrix(mat, row, col)
	return MatrixDeterminant(submatrix)
}

func MatrixCofactor(mat Matrix, row, col int) float64 {
	cofactor := MatrixMinor(mat, row, col)

	if (row+col)%2 != 0 {
		cofactor *= -1
	}

	return cofactor
}

func MatrixDeterminant(mat Matrix) float64 {
	if mat.Rows != mat.Cols {
		panic("The determinant is only defined for square matrices")
	}

	if mat.Cols == 2 {
		return mat.Get(0, 0)*mat.Get(1, 1) - mat.Get(0, 1)*mat.Get(1, 0)
	}

	var determinant float64 = 0.0

	for col := 0; col < mat.Cols; col++ {
		determinant += (mat.Get(0, col) * MatrixCofactor(mat, 0, col))
	}

	return determinant
}

func MatrixIsInvertible(mat Matrix) bool {
	determinant := MatrixDeterminant(mat)

	return !(determinant == 0)
}

func MatrixInverse(mat Matrix) Matrix {
	if mat.Rows != mat.Cols {
		panic("The inverse is only defined for square matrices")
	}

	if !MatrixIsInvertible(mat) {
		panic("The given matrix is not invertible")
	}

	inverse := MakeMatrix(make([]float64, mat.Rows*mat.Cols), mat.Rows, mat.Cols)
	determinant := MatrixDeterminant(mat)

	for row := 0; row < mat.Rows; row++ {
		for col := 0; col < mat.Cols; col++ {
			cofactor := MatrixCofactor(mat, row, col)
			inverse.Set(cofactor/determinant, col, row)
		}
	}

	return inverse
}

func (mat Matrix) Get(row, col int) float64 {
	return mat.elems[row*mat.Cols+col]
}

func (mat *Matrix) Set(val float64, row, col int) {
	mat.elems[row*mat.Cols+col] = val
}

func (mat Matrix) GetByIndex(index int) float64 {
	return mat.elems[index]
}

func multiplyRowCol(mat1, mat2 Matrix, row, col int) float64 {
	result := 0.0

	for i := 0; i < mat1.Rows; i++ {
		result += mat1.Get(row, i) * mat2.Get(i, col)
	}

	return result
}
