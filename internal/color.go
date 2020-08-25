package internal

import "github.com/google/go-cmp/cmp"

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{r, g, b}
}

func AddColors(c1, c2 Color) Color {
	return NewColor(c1.R+c2.R, c1.G+c2.G, c1.B+c2.B)
}

func (c1 Color) Add(c2 Color) Color {
	return NewColor(c1.R+c2.R, c1.G+c2.G, c1.B+c2.B)
}

func SubColors(c1, c2 Color) Color {
	return NewColor(c1.R-c2.R, c1.G-c2.G, c1.B-c2.B)
}

func (c1 Color) Sub(c2 Color) Color {
	return NewColor(c1.R-c2.R, c1.G-c2.G, c1.B-c2.B)
}

func ColorScalarMultiply(c1 Color, scalar float64) Color {
	return NewColor(c1.R*scalar, c1.G*scalar, c1.B*scalar)
}

func (c1 Color) MultiplyByScalar(scalar float64) Color {
	return NewColor(c1.R*scalar, c1.G*scalar, c1.B*scalar)
}

func ColorEquals(t1, t2 Color) bool {
	return cmp.Equal(t1, t2, opt)
}

func (t1 Color) Equals(t2 Color) bool {
	return cmp.Equal(t1, t2, opt)
}

func HadamardProduct(c1, c2 Color) Color {
	return NewColor(
		c1.R*c2.R,
		c1.G*c2.G,
		c1.B*c2.B,
	)
}

func (c1 Color) HadamardProduct(c2 Color) Color {
	return NewColor(
		c1.R*c2.R,
		c1.G*c2.G,
		c1.B*c2.B,
	)
}
