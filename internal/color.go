package internal

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
