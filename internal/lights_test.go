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
	result := Lighting(m, NewSphere(), light, position, eyeV, normalV, 1.0)

	assert.True(t, ColorEquals(NewColor(1.9, 1.9, 1.9), result))
}

func TestLightingEyeBetweenLightAndSurfaceLightOffset45(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 1/math.Sqrt(2), -1/math.Sqrt(2))
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result := Lighting(m, NewSphere(), light, position, eyeV, normalV, 1.0)

	assert.True(t, ColorEquals(NewColor(1.0, 1.0, 1.0), result))
}

func TestLightingEyeOppositeSurfaceLightOffset45(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	result := Lighting(m, NewSphere(), light, position, eyeV, normalV, 1.0)

	assert.True(t, ColorEquals(NewColor(0.7364, 0.7364, 0.7364), result))
}

func TestLightingEyeInPathOfReflectionVector(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, -1/math.Sqrt(2), -1/math.Sqrt(2))
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	result := Lighting(m, NewSphere(), light, position, eyeV, normalV, 1.0)

	assert.True(t, ColorEquals(NewColor(1.6364, 1.6364, 1.6364), result))
}

func TestLightingLightBehindSurface(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), NewColor(1, 1, 1))
	result := Lighting(m, NewSphere(), light, position, eyeV, normalV, 1.0)

	assert.True(t, ColorEquals(NewColor(0.1, 0.1, 0.1), result))
}

func LightingWithSurfaceInShadow(t *testing.T) {
	m := NewDefaultMaterial()
	position := NewPoint(0, 0, 0)

	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	result := Lighting(m, NewSphere(), light, position, eyeV, normalV, 0.0)

	assert.True(t, ColorEquals(NewColor(0.1, 0.1, 0.1), result))
}

func TestLightingWithPatternApplied(t *testing.T) {
	m := NewDefaultMaterial()

	m.Pattern = NewStripePattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	m.Ambient = 1
	m.Diffuse = 0
	m.Specular = 0
	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	c1 := Lighting(m, NewSphere(), light, NewPoint(0.9, 0, 0), eyeV, normalV, 1.0)
	c2 := Lighting(m, NewSphere(), light, NewPoint(1.1, 0, 0), eyeV, normalV, 1.0)

	assert.True(t, ColorEquals(NewColor(1, 1, 1), c1))
	assert.True(t, ColorEquals(NewColor(0, 0, 0), c2))
}

func TestPointLightsEvaluateIntensityAtPoint(t *testing.T) {
	testCases := []struct {
		point  Tuple
		result float64
	}{
		{NewPoint(0, 1.0001, 0), 1.0},
		{NewPoint(-1.0001, 0, 0), 1.0},
		{NewPoint(0, 0, -1.0001), 1.0},
		{NewPoint(0, 0, 1.0001), 0.0},
		{NewPoint(1.0001, 0, 0), 0.0},
		{NewPoint(1.0001, 0, 0), 0.0},
		{NewPoint(0, 0, 0), 0.0},
	}

	w := NewDefaultWorld()
	light := w.Lights[0]

	for _, test := range testCases {
		intensity := IntensityAt(light, test.point, w)
		assert.InDelta(t, test.result, intensity, float64EqualityThreshold)
	}
}

func TestLightingUsesIntensityToAttenuateColor(t *testing.T) {
	testCases := []struct {
		intensity float64
		result    Color
	}{
		{1.0, NewColor(1, 1, 1)},
		{0.5, NewColor(0.55, 0.55, 0.55)},
		{0.0, NewColor(0.1, 0.1, 0.1)},
	}

	w := NewDefaultWorld()
	w.Lights[0] = NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))

	shape := w.Objects[0]
	shapeMaterial := shape.GetMaterial()
	shapeMaterial.Ambient = 0.1
	shapeMaterial.Diffuse = 0.9
	shapeMaterial.Specular = 0.0
	shapeMaterial.SetColor(NewColor(1, 1, 1))
	shape.SetMaterial(shapeMaterial)

	pt := NewPoint(0, 0, -1)
	eyeV := NewVector(0, 0, -1)
	normalV := NewVector(0, 0, -1)

	for _, test := range testCases {
		result := Lighting(shape.GetMaterial(), shape, w.Lights[0], pt, eyeV, normalV, test.intensity)
		assert.True(t, ColorEquals(test.result, result))
	}
}

func TestCreateAreaLight(t *testing.T) {
	corner := NewPoint(0, 0, 0)
	v1 := NewVector(2, 0, 0)
	v2 := NewVector(0, 0, 1)
	light := NewAreaLight(corner, v1, 4, v2, 2, NewColor(1, 1, 1))

	assert.True(t, TupleEquals(corner, light.CornerPos))
	assert.Equal(t, 4, light.USteps)
	assert.True(t, TupleEquals(NewVector(0.5, 0, 0), light.UVec))
	assert.Equal(t, 2, light.VSteps)
	assert.True(t, TupleEquals(NewVector(0, 0, 0.5), light.VVec))
	assert.Equal(t, 8, light.Samples)
	assert.True(t, TupleEquals(NewPoint(1.0, 0.0, 0.5), light.Position))
}

func TestFindSinglePointOnAreaLight(t *testing.T) {
	testCases := []struct {
		u, v   int
		result Tuple
	}{
		{0, 0, NewPoint(0.25, 0, 0.25)},
		{1, 0, NewPoint(0.75, 0, 0.25)},
		{0, 1, NewPoint(0.25, 0, 0.75)},
		{2, 0, NewPoint(1.25, 0, 0.25)},
		{3, 1, NewPoint(1.75, 0, 0.75)},
	}

	corner := NewPoint(0, 0, 0)
	v1 := NewVector(2, 0, 0)
	v2 := NewVector(0, 0, 1)
	light := NewAreaLight(corner, v1, 4, v2, 2, NewColor(1, 1, 1))

	for _, test := range testCases {
		pt := light.PointOnLight(test.u, test.v, false)
		assert.True(t, TupleEquals(test.result, pt))
	}
}

func TestAreaLightIntensityFunction(t *testing.T) {
	testCases := []struct {
		point  Tuple
		result float64
	}{
		{NewPoint(0, 0, 2), 0.0},
		{NewPoint(1, -1, 2), 0.25},
		{NewPoint(1.5, 0, 2), 0.5},
		{NewPoint(1.25, 1.25, 3), 0.75},
		{NewPoint(0, 0, -2), 1.0},
	}

	w := NewDefaultWorld()
	corner := NewPoint(-0.5, -0.5, -5)
	v1 := NewVector(1, 0, 0)
	v2 := NewVector(0, 1, 0)
	light := NewAreaLight(corner, v1, 2, v2, 2, NewColor(1, 1, 1))

	for _, test := range testCases {
		intensity := IntensityAt(light, test.point, w)
		assert.InDelta(t, test.result, intensity, float64EqualityThreshold)
	}
}

func TestLightingFunctionSamplesAreaLight(t *testing.T) {
	testCases := []struct {
		point  Tuple
		result Color
	}{
		{NewPoint(0, 0, -1), NewColor(0.9965, 0.9965, 0.9965)},
		{NewPoint(0, 0.7071, -0.7071), NewColor(0.6232, 0.6232, 0.6232)},
	}

	corner := NewPoint(-0.5, -0.5, -5)
	v1 := NewVector(1, 0, 0)
	v2 := NewVector(0, 1, 0)

	light := NewAreaLight(corner, v1, 2, v2, 2, NewColor(1, 1, 1))

	shape := NewSphere()
	shapeMaterial := shape.GetMaterial()
	shapeMaterial.Ambient = 0.1
	shapeMaterial.Diffuse = 0.9
	shapeMaterial.Specular = 0.0
	shapeMaterial.SetColor(NewColor(1, 1, 1))
	shape.SetMaterial(shapeMaterial)

	eye := NewPoint(0, 0, -5)

	for _, test := range testCases {
		pt := test.point
		eyeV := Normalize(SubTuples(eye, pt))
		normalV := NewVector(pt.X, pt.Y, pt.Z)
		result := Lighting(shape.GetMaterial(), shape, light, pt, eyeV, normalV, 1.0)

		assert.True(t, ColorEquals(test.result, result))
	}
}
