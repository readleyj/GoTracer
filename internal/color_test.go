package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	assert.InDelta(t, -0.5, c.R, float64EqualityThreshold)
	assert.InDelta(t, 0.4, c.G, float64EqualityThreshold)
	assert.InDelta(t, 1.7, c.B, float64EqualityThreshold)
}

func TestColorAdd(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	c3 := AddColors(c1, c2)

	assert.InDelta(t, 1.6, c3.R, float64EqualityThreshold)
	assert.InDelta(t, 0.7, c3.G, float64EqualityThreshold)
	assert.InDelta(t, 1.0, c3.B, float64EqualityThreshold)
}

func TestColorSub(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	c3 := SubColors(c1, c2)

	assert.InDelta(t, 0.2, c3.R, float64EqualityThreshold)
	assert.InDelta(t, 0.5, c3.G, float64EqualityThreshold)
	assert.InDelta(t, 0.5, c3.B, float64EqualityThreshold)
}

func TestMultiplyColorByScalar(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)
	c2 := ColorScalarMultiply(c1, 2.0)

	assert.InDelta(t, 0.4, c2.R, float64EqualityThreshold)
	assert.InDelta(t, 0.6, c2.G, float64EqualityThreshold)
	assert.InDelta(t, 0.8, c2.B, float64EqualityThreshold)
}

func TestHadamardProduct(t *testing.T) {
	c1 := NewColor(1.0, 0.2, 0.4)
	c2 := NewColor(0.9, 1.0, 0.1)
	c3 := HadamardProduct(c1, c2)

	assert.InDelta(t, 0.9, c3.R, float64EqualityThreshold)
	assert.InDelta(t, 0.2, c3.G, float64EqualityThreshold)
	assert.InDelta(t, 0.04, c3.B, float64EqualityThreshold)
}
