package internal

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructCamera(t *testing.T) {
	hsize, vsize := 160, 120
	fov := math.Pi / 2
	c := NewCamera(hsize, vsize, fov)

	assert.Equal(t, 160, c.Hsize)
	assert.Equal(t, 120, c.Vsize)
	assert.InDelta(t, math.Pi/2, c.FOV, float64EqualityThreshold)
	assert.True(t, MatrixEquals(Identity4, c.Transform))
}

func TestHorizontalCanvasPixelSize(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)

	assert.InDelta(t, 0.01, c.PixelSize, float64EqualityThreshold)
}

func TestVerticalCanvasPixelSize(t *testing.T) {
	c := NewCamera(125.0, 200.0, math.Pi/2)

	assert.InDelta(t, 0.01, c.PixelSize, float64EqualityThreshold)
}

func TestRayThroughCanvasCenter(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := RayForPixel(c, 100, 50)

	assert.True(t, TupleEquals(NewPoint(0, 0, 0), r.Origin))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), r.Direction))
}

func TestRayThroughCanvasCorner(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	r := RayForPixel(c, 0, 0)

	assert.True(t, TupleEquals(NewPoint(0, 0, 0), r.Origin))
	assert.True(t, TupleEquals(NewVector(0.66519, 0.33259, -0.66851), r.Direction))
}

func TestRayWithTransformedCamera(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.Transform = MatrixMultiply(RotateY(math.Pi/4), Translate(0, -2, 5))
	r := RayForPixel(c, 100, 50)

	assert.True(t, TupleEquals(NewPoint(0, 2, -5), r.Origin))
	assert.True(t, TupleEquals(NewVector(1/math.Sqrt(2), 0, -1/math.Sqrt(2)), r.Direction))
}

func TestRenderingWorldWithCamera(t *testing.T) {
	w := NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)

	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	c.Transform = ViewTransform(from, to, up)

	image := Render(c, w)
	pixelColor := image.GetColorAtPixel(5, 5)

	assert.InDelta(t, 0.38066, pixelColor.R, float64EqualityThreshold)
	assert.InDelta(t, 0.47583, pixelColor.G, float64EqualityThreshold)
	assert.InDelta(t, 0.28550, pixelColor.B, float64EqualityThreshold)
}
