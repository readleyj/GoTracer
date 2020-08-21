package internal

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanvasCreation(t *testing.T) {
	canvas := CreateCanvas(10, 20)

	for _, pixel := range canvas.Pixels {
		assert.InDelta(t, 0.0, pixel.R, float64EqualityThreshold)
		assert.InDelta(t, 0.0, pixel.G, float64EqualityThreshold)
		assert.InDelta(t, 0.0, pixel.B, float64EqualityThreshold)
	}
}

func TestWritePixel(t *testing.T) {
	canvas := CreateCanvas(10, 20)
	red := Color{1, 0, 0}

	canvas.WritePixelAtCoord(2, 3, red)
	pixel := canvas.Pixels[canvas.GetPixelIndex(2, 3)]

	assert.InDelta(t, 1.0, pixel.R, float64EqualityThreshold)
	assert.InDelta(t, 0.0, pixel.G, float64EqualityThreshold)
	assert.InDelta(t, 0.0, pixel.B, float64EqualityThreshold)
}

func TestPpmHeader(t *testing.T) {
	canvas := CreateCanvas(5, 3)

	ppmString := canvas.ToPPM()
	header := `P3
5 3
255
`
	assert.True(t, strings.HasPrefix(ppmString, header))
}

func TestPpmPixels(t *testing.T) {
	canvas := CreateCanvas(5, 3)
	c1 := Color{1.5, 0, 0}
	c2 := Color{0.0, 0.5, 0}
	c3 := Color{-0.5, 0.0, 1.0}

	canvas.WritePixelAtCoord(0, 0, c1)
	canvas.WritePixelAtCoord(2, 1, c2)
	canvas.WritePixelAtCoord(4, 2, c3)

	ppmString := canvas.ToPPM()
	ppmPixels := strings.Join(strings.Split(ppmString, "\n")[3:], "\n")
	pixels := `255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`
	assert.Equal(t, pixels, ppmPixels)
}

func TestPpmStringLength(t *testing.T) {
	canvas := CreateCanvas(10, 2)
	color := Color{1, 0.8, 0.6}

	for i := range canvas.Pixels {
		canvas.WritePixelAtIndex(i, color)
	}

	ppmString := canvas.ToPPM()
	ppmPixels := strings.Join(strings.Split(ppmString, "\n")[3:], "\n")
	pixels := `255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`
	assert.Equal(t, pixels, ppmPixels)
}

func TestNewlineSuffix(t *testing.T) {
	canvas := CreateCanvas(5, 3)
	ppmString := canvas.ToPPM()

	fmt.Print(ppmString[len(ppmString)-1])

	assert.Equal(t, "\n", string(ppmString[len(ppmString)-1]))
}
