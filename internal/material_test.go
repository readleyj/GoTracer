package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultMaterial(t *testing.T) {
	m := NewDefaultMaterial()

	assert.True(t, ColorEquals(NewColor(1, 1, 1), m.Color))
	assert.InDelta(t, 0.1, m.Ambient, float64EqualityThreshold)
	assert.InDelta(t, 0.9, m.Diffuse, float64EqualityThreshold)
	assert.InDelta(t, 0.9, m.Specular, float64EqualityThreshold)
	assert.InDelta(t, 200.0, m.Shininess, float64EqualityThreshold)
	assert.InDelta(t, 0.0, m.Reflective, float64EqualityThreshold)
	assert.InDelta(t, 0.0, m.Transparency, float64EqualityThreshold)
	assert.InDelta(t, 1.0, m.RefractiveIndex, float64EqualityThreshold)
}
