package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSphereDefaultTransform(t *testing.T) {
	s := NewSphere()
	assert.True(t, MatrixEquals(s.Transform, Identity4))
}

func TestSphereSetTransform(t *testing.T) {
	s := NewSphere()
	transform := Translate(2, 3, 4)
	s.SetTransform(transform)
	assert.True(t, MatrixEquals(s.Transform, transform))
}

func TestSphereHasDefaultMaterial(t *testing.T) {
	s := NewSphere()
	m := s.Material
	target := NewDefaultMaterial()

	assert.True(t, ColorEquals(target.Color, m.Color))
	assert.InDelta(t, target.Ambient, m.Ambient, float64EqualityThreshold)
	assert.InDelta(t, target.Diffuse, m.Diffuse, float64EqualityThreshold)
	assert.InDelta(t, target.Specular, m.Specular, float64EqualityThreshold)
	assert.InDelta(t, target.Shininess, m.Shininess, float64EqualityThreshold)
}

func TestSphereCanBeAssignedMaterial(t *testing.T) {
	s := NewSphere()
	m := NewMaterial(
		DefaultMaterial.Color,
		1.0,
		DefaultMaterial.Diffuse,
		DefaultMaterial.Specular,
		DefaultMaterial.Shininess,
	)
	s.SetMaterial(m)
	sphereMaterial := s.Material

	assert.True(t, ColorEquals(m.Color, sphereMaterial.Color))
	assert.InDelta(t, m.Ambient, sphereMaterial.Ambient, float64EqualityThreshold)
	assert.InDelta(t, m.Diffuse, sphereMaterial.Diffuse, float64EqualityThreshold)
	assert.InDelta(t, m.Specular, sphereMaterial.Specular, float64EqualityThreshold)
	assert.InDelta(t, m.Shininess, sphereMaterial.Shininess, float64EqualityThreshold)
}
