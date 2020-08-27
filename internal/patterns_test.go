package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStripePattern(t *testing.T) {
	pattern := NewStripePattern(white, black)

	assert.True(t, ColorEquals(white, pattern.A))
	assert.True(t, ColorEquals(black, pattern.B))
}

func TestStripePatternIsConstantInY(t *testing.T) {
	pattern := NewStripePattern(white, black)

	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 0, 0))))
	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 1, 0))))
	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 2, 0))))
}

func TestStripePatternIsConstantInZ(t *testing.T) {
	pattern := NewStripePattern(white, black)

	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 0, 0))))
	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 0, 1))))
	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 0, 2))))
}

func TestStripePatternAlternatesInX(t *testing.T) {
	pattern := NewStripePattern(white, black)

	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0, 0, 0))))
	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(0.9, 0, 1))))
	assert.True(t, ColorEquals(black, pattern.PatternAt(NewPoint(1, 0, 0))))
	assert.True(t, ColorEquals(black, pattern.PatternAt(NewPoint(-0.1, 0, 0))))
	assert.True(t, ColorEquals(black, pattern.PatternAt(NewPoint(-1, 0, 1))))
	assert.True(t, ColorEquals(white, pattern.PatternAt(NewPoint(-1.1, 0, 0))))
}

func TestStripesWithObjectTransform(t *testing.T) {
	object := NewSphere()
	object.SetTransform(Scale(2, 2, 2))
	pattern := NewStripePattern(white, black)
	c := PatternAtShape(pattern, object, NewPoint(1.5, 0, 0))

	assert.True(t, ColorEquals(white, c))
}

func TestStripesWithPatternTransform(t *testing.T) {
	object := NewSphere()
	pattern := NewStripePattern(white, black)
	pattern.SetTransform(Scale(2, 2, 2))
	c := PatternAtShape(pattern, object, NewPoint(1.5, 0, 0))

	assert.True(t, ColorEquals(white, c))
}

func TestStripesWithObjectAndPatternTransform(t *testing.T) {
	object := NewSphere()
	object.SetTransform(Scale(2, 2, 2))
	pattern := NewStripePattern(white, black)
	pattern.SetTransform(Translate(0.5, 0, 0))
	c := PatternAtShape(pattern, object, NewPoint(2.5, 0, 0))

	assert.True(t, ColorEquals(white, c))
}

func TestDefaultPatternTransform(t *testing.T) {
	pattern := NewTestPattern()

	assert.True(t, MatrixEquals(Identity4, pattern.Transform))
}

func TestAssignPatternTransform(t *testing.T) {
	pattern := NewTestPattern()
	pattern.SetTransform(Translate(1, 2, 3))

	assert.True(t, MatrixEquals(Translate(1, 2, 3), pattern.Transform))
}

func TestPatternWithObjectTransform(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(Scale(2, 2, 2))
	pattern := NewTestPattern()
	c := PatternAtShape(pattern, shape, NewPoint(2, 3, 4))

	assert.True(t, ColorEquals(NewColor(1, 1.5, 2), c))
}

func TestPatternWithPatternTransform(t *testing.T) {
	shape := NewSphere()
	pattern := NewTestPattern()
	pattern.SetTransform(Scale(2, 2, 2))
	c := PatternAtShape(pattern, shape, NewPoint(2, 3, 4))

	assert.True(t, ColorEquals(NewColor(1, 1.5, 2), c))
}

func TestPatternWithObjectAndPatternTransform(t *testing.T) {
	shape := NewSphere()
	shape.SetTransform(Scale(2, 2, 2))
	pattern := NewTestPattern()
	pattern.SetTransform(Translate(0.5, 1, 1.5))
	c := PatternAtShape(pattern, shape, NewPoint(2.5, 3, 3.5))

	assert.True(t, ColorEquals(NewColor(0.75, 0.5, 0.25), c))
}

func TestGradientLinearlyInterpolatesBetweenColors(t *testing.T) {
	pattern := NewGradientPattern(white, black)

	color1 := pattern.PatternAt(NewPoint(0, 0, 0))
	color2 := pattern.PatternAt(NewPoint(0.25, 0, 0))
	color3 := pattern.PatternAt(NewPoint(0.5, 0, 0))
	color4 := pattern.PatternAt(NewPoint(0.75, 0, 0))

	assert.True(t, ColorEquals(white, color1))
	assert.True(t, ColorEquals(NewColor(0.75, 0.75, 0.75), color2))
	assert.True(t, ColorEquals(NewColor(0.5, 0.5, 0.5), color3))
	assert.True(t, ColorEquals(NewColor(0.25, 0.25, 0.25), color4))
}

func TestRingShouldExtendInXZ(t *testing.T) {
	pattern := NewRingPattern(white, black)

	color1 := pattern.PatternAt(NewPoint(0, 0, 0))
	color2 := pattern.PatternAt(NewPoint(1, 0, 0))
	color3 := pattern.PatternAt(NewPoint(0, 0, 1))
	color4 := pattern.PatternAt(NewPoint(0.708, 0, 0.708))

	assert.True(t, ColorEquals(white, color1))
	assert.True(t, ColorEquals(black, color2))
	assert.True(t, ColorEquals(black, color3))
	assert.True(t, ColorEquals(black, color4))
}

func TestCheckersShouldRepeatInX(t *testing.T) {
	pattern := NewCheckersPattern(white, black)

	color1 := pattern.PatternAt(NewPoint(0, 0, 0))
	color2 := pattern.PatternAt(NewPoint(0.99, 0, 0))
	color3 := pattern.PatternAt(NewPoint(1.01, 0, 0))

	assert.True(t, ColorEquals(white, color1))
	assert.True(t, ColorEquals(white, color2))
	assert.True(t, ColorEquals(black, color3))
}

func TestCheckersShouldRepeatInY(t *testing.T) {
	pattern := NewCheckersPattern(white, black)

	color1 := pattern.PatternAt(NewPoint(0, 0, 0))
	color2 := pattern.PatternAt(NewPoint(0, 0.99, 0))
	color3 := pattern.PatternAt(NewPoint(0, 1.01, 0))

	assert.True(t, ColorEquals(white, color1))
	assert.True(t, ColorEquals(white, color2))
	assert.True(t, ColorEquals(black, color3))
}

func TestCheckersShouldRepeatInZ(t *testing.T) {
	pattern := NewCheckersPattern(white, black)

	color1 := pattern.PatternAt(NewPoint(0, 0, 0))
	color2 := pattern.PatternAt(NewPoint(0, 0, 0.99))
	color3 := pattern.PatternAt(NewPoint(0, 0, 1.01))

	assert.True(t, ColorEquals(white, color1))
	assert.True(t, ColorEquals(white, color2))
	assert.True(t, ColorEquals(black, color3))
}
