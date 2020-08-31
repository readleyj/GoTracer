package internal

import "math"

var black Color = NewColor(0, 0, 0)
var white Color = NewColor(1, 1, 1)

type Pattern interface {
	PatternAt(point Tuple) Color
	SetTransform(transform Matrix)
	GetTransform() Matrix
	GetInverse() Matrix
}

func PatternAtShape(pattern Pattern, shape Shape, worldPoint Tuple) Color {
	objectPoint := WorldToObject(shape, worldPoint)
	patternPoint := MatrixTupleMultiply(pattern.GetInverse(), objectPoint)

	return pattern.PatternAt(patternPoint)
}

type TestPattern struct {
	A         Color
	B         Color
	Transform Matrix
	Inverse   Matrix
}

func NewTestPattern() *TestPattern {
	return &TestPattern{
		A:         white,
		B:         black,
		Transform: NewIdentity4(),
		Inverse:   NewIdentity4(),
	}
}

func (p *TestPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *TestPattern) SetTransform(transform Matrix) {
	p.Transform = transform
	p.Inverse = MatrixInverse(p.Transform)
}

func (p *TestPattern) GetInverse() Matrix {
	return p.Inverse
}

func (p *TestPattern) PatternAt(point Tuple) Color {
	return NewColor(point.X, point.Y, point.Z)
}

type StripePattern struct {
	A         Color
	B         Color
	Transform Matrix
	Inverse   Matrix
}

func NewStripePattern(a, b Color) *StripePattern {
	return &StripePattern{
		A:         a,
		B:         b,
		Transform: NewIdentity4(),
		Inverse:   NewIdentity4(),
	}
}

func (p *StripePattern) PatternAt(point Tuple) Color {
	if int(math.Floor(point.X))%2 == 0 {
		return p.A
	}

	return p.B
}

func (p *StripePattern) SetTransform(transform Matrix) {
	p.Transform = transform
	p.Inverse = MatrixInverse(p.Transform)
}

func (p *StripePattern) GetTransform() Matrix {
	return p.Transform
}

func (p *StripePattern) GetInverse() Matrix {
	return p.Inverse
}

type GradientPattern struct {
	A         Color
	B         Color
	Transform Matrix
	Inverse   Matrix
}

func NewGradientPattern(a, b Color) *GradientPattern {
	return &GradientPattern{
		A:         a,
		B:         b,
		Transform: NewIdentity4(),
		Inverse:   NewIdentity4(),
	}
}

func (p *GradientPattern) PatternAt(point Tuple) Color {
	distance := SubColors(p.B, p.A)
	fraction := point.X - math.Floor(point.X)

	return AddColors(p.A, ColorScalarMultiply(distance, fraction))
}

func (p *GradientPattern) SetTransform(transform Matrix) {
	p.Transform = transform
	p.Inverse = MatrixInverse(p.Transform)
}

func (p *GradientPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *GradientPattern) GetInverse() Matrix {
	return p.Inverse
}

type RingPattern struct {
	A         Color
	B         Color
	Transform Matrix
	Inverse   Matrix
}

func NewRingPattern(a, b Color) *RingPattern {
	return &RingPattern{
		A:         a,
		B:         b,
		Transform: NewIdentity4(),
		Inverse:   NewIdentity4(),
	}
}

func (p *RingPattern) PatternAt(point Tuple) Color {
	value := int(math.Floor(math.Sqrt(point.X*point.X + point.Z*point.Z)))

	if value%2 == 0 {
		return p.A
	}

	return p.B
}

func (p *RingPattern) SetTransform(transform Matrix) {
	p.Transform = transform
	p.Inverse = MatrixInverse(p.Transform)
}

func (p *RingPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *RingPattern) GetInverse() Matrix {
	return p.Inverse
}

type CheckersPattern struct {
	A         Color
	B         Color
	Transform Matrix
	Inverse   Matrix
}

func NewCheckersPattern(a, b Color) *CheckersPattern {
	return &CheckersPattern{
		A:         a,
		B:         b,
		Transform: NewIdentity4(),
		Inverse:   NewIdentity4(),
	}
}

func (p *CheckersPattern) PatternAt(point Tuple) Color {
	sum := int(math.Floor(point.X) + math.Floor(point.Y) + math.Floor(point.Z))

	if sum%2 == 0 {
		return p.A
	}

	return p.B
}

func (p *CheckersPattern) SetTransform(transform Matrix) {
	p.Transform = transform
	p.Inverse = MatrixInverse(p.Transform)
}

func (p *CheckersPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *CheckersPattern) GetInverse() Matrix {
	return p.Inverse
}
