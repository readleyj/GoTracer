package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatingWorld(t *testing.T) {
	w := NewWorld()
	emptyLight := PointLight{}

	assert.True(t, w.Light == emptyLight)
	assert.Equal(t, 0, len(w.Objects))
}

func TestDefaultWorld(t *testing.T) {
	light := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))

	s1 := NewSphere()
	s1.SetMaterial(NewMaterial(
		NewColor(0.8, 1.0, 0.6),
		DefaultMaterial.Ambient,
		0.7,
		0.2,
		DefaultMaterial.Shininess,
	))

	s2 := NewSphere()
	s2.SetTransform(Scale(0.5, 0.5, 0.5))

	w := NewDefaultWorld()
	assert.True(t, PointLightEquals(light, w.Light))
	assert.True(t, w.ContainsObject(s1))
	assert.True(t, w.ContainsObject(s2))
}

func TestIntersectWorldWithRay(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	xs := IntersectWorld(w, r)

	assert.Equal(t, 4, len(xs))
	assert.InDelta(t, 4.0, xs[0].T, float64EqualityThreshold)
	assert.InDelta(t, 4.5, xs[1].T, float64EqualityThreshold)
	assert.InDelta(t, 5.5, xs[2].T, float64EqualityThreshold)
	assert.InDelta(t, 6.0, xs[3].T, float64EqualityThreshold)
}
