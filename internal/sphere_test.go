package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGlassSphere(t *testing.T) {
	s := NewGlassSphere()

	assert.True(t, MatrixEquals(Identity4, s.GetTransform()))
	assert.InDelta(t, 1.0, s.GetMaterial().Transparency, float64EqualityThreshold)
	assert.InDelta(t, 1.5, s.GetMaterial().RefractiveIndex, float64EqualityThreshold)
}
