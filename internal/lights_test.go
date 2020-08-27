package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointLightHasPositionAndIntensity(t *testing.T) {
	intensity := NewColor(1, 1, 1)
	position := NewPoint(0, 0, 0)
	light := NewPointLight(position, intensity)

	assert.True(t, TupleEquals(position, light.Position))
	assert.True(t, ColorEquals(intensity, light.Intensity))
}

func TestLightingEyeBetweenLightAndSurface(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result := Lighting(m, light, position, eyeV, normalV, false)

	assert.True(t, ColorEquals(NewColor(1.9, 1.9, 1.9), result))
}

func TestLightingEyeBetweenLightAndSurfaceLightOffset45(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 1/math.Sqrt(2), -1/math.Sqrt(2))
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result := Lighting(m, light, position, eyeV, normalV, false)

	assert.True(t, ColorEquals(NewColor(1.0, 1.0, 1.0), result))
}

func TestLightingEyeOppositeSurfaceLightOffset45(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	result := Lighting(m, light, position, eyeV, normalV, false)

	assert.True(t, ColorEquals(NewColor(0.7364, 0.7364, 0.7364), result))
}

func TestLightingEyeInPathOfReflectionVector(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, -1/math.Sqrt(2), -1/math.Sqrt(2))
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	result := Lighting(m, light, position, eyeV, normalV, false)

	assert.True(t, ColorEquals(NewColor(1.6364, 1.6364, 1.6364), result))
}

func TestLightingLightBehindSurface(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), NewColor(1, 1, 1))
	result := Lighting(m, light, position, eyeV, normalV, false)

	assert.True(t, ColorEquals(NewColor(0.1, 0.1, 0.1), result))
}

func LightingWithSurfaceInShadow(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	inShadow := true
	result := Lighting(m, light, position, eyeV, normalV, inShadow)

	assert.True(t, ColorEquals(NewColor(0.1, 0.1, 0.1), result))
}
