package internal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Canvas struct {
	W      int
	H      int
	Pixels []Color
}

func NewCanvas(W, H int) *Canvas {
	pixels := make([]Color, W*H)

	for i := range pixels {
		pixels[i] = NewColor(0, 0, 0)
	}

	return &Canvas{
		W:      W,
		H:      H,
		Pixels: pixels,
	}
}

func (c *Canvas) WritePixelAtCoord(x, y int, col Color) {
	if x < 0 || y < 0 || x >= c.W || y >= c.H {
		fmt.Println("Pixel coordinates are out of bounds")
		return
	}

	writeIndex := c.GetPixelIndex(x, y)
	c.Pixels[writeIndex] = col
}

func (c *Canvas) WritePixelAtIndex(index int, col Color) {
	if index > c.GetLastIndex() {
		fmt.Println("Pixel index is out of bounds")
		return
	}

	c.Pixels[index] = col
}

func (c *Canvas) GetLastIndex() int {
	return c.W*c.H - 1
}

func (c *Canvas) GetPixelIndex(x, y int) int {
	return x + c.W*y
}

func (c *Canvas) GetColorAtPixel(x, y int) Color {
	index := c.GetPixelIndex(x, y)
	return c.Pixels[index]
}

func (c *Canvas) ToPPM() string {
	ppm := fmt.Sprintf("P3\n%d %d\n255\n", c.W, c.H)

	for row := 0; row < c.H; row++ {
		var b strings.Builder
		writtenLen := 0

		for col := 0; col < c.W; col++ {
			color := c.GetColorAtPixel(col, row)

			formatPPM(color.R, &b, &writtenLen)
			formatPPM(color.G, &b, &writtenLen)
			formatPPM(color.B, &b, &writtenLen)
		}

		res := strings.TrimSuffix(b.String(), " ")
		ppm += res + "\n"
	}

	return ppm
}

func formatPPM(colorValue float64, b *strings.Builder, writtenLen *int) {
	clamped := clamp(colorValue)
	clampedStr := strconv.Itoa(clamped)

	if *writtenLen+len(clampedStr)+2 > 70 {
		b.WriteString("\n")
		b.WriteString(clampedStr)
		*writtenLen = len(clampedStr)
	} else {
		addLen := 1

		if *writtenLen == 0 {
			addLen = 0
		} else {
			b.WriteString(" ")
		}

		b.WriteString(clampedStr)
		*writtenLen += len(clampedStr) + addLen
	}
}

func clamp(pixel float64) int {
	scaled := math.Ceil(pixel * 255.0)
	scaled = math.Max(0.0, scaled)
	scaled = math.Min(255.0, scaled)

	return int(scaled)
}
