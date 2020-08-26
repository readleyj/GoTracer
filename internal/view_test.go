package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultOrientationTransform(t *testing.T) {
	from := NewPoint(0, 0, 0)
	to := NewPoint(0, 0, -1)
	up := NewVector(0, 1, 0)
	transform := ViewTransform(from, to, up)

	assert.True(t, MatrixEquals(Identity4, transform))
}

func TestViewTransformLookingInPositiveZ(t *testing.T) {
	from := NewPoint(0, 0, 8)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	transform := ViewTransform(from, to, up)

	assert.True(t, MatrixEquals(Translate(0, 0, -8), transform))
}

func TestArbitraryViewTransform(t *testing.T) {
	from := NewPoint(1, 3, 2)
	to := NewPoint(4, -2, 8)
	up := NewVector(1, 1, 0)
	transform := ViewTransform(from, to, up)

	target := NewMatrix4([]float64{
		-0.50709, 0.50709, 0.67612, -2.36643,
		0.76772, 0.60609, 0.12122, -2.82843,
		-0.35857, 0.59761, -0.71714, 0.0,
		0.0, 0.0, 0.0, 1.0,
	})

	assert.True(t, MatrixEquals(target, transform))
}
