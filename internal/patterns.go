package internal

import "math"

var black Color = NewColor(0, 0, 0)
var white Color = NewColor(1, 1, 1)

type Pattern interface {
	SetTransform(transform Matrix)
	GetTransform() Matrix
	PatternAt(point Tuple) Color
}

func PatternAtShape(pattern Pattern, shape Shape, worldPoint Tuple) Color {
	objectPoint := MatrixTupleMultiply(MatrixInverse(shape.GetTransform()), worldPoint)
	patternPoint := MatrixTupleMultiply(MatrixInverse(pattern.GetTransform()), objectPoint)

	return pattern.PatternAt(patternPoint)
}

type TestPattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NewTestPattern() *TestPattern {
	return &TestPattern{white, black, NewIdentity4()}
}

func (p *TestPattern) GetTransform() Matrix {
	return p.Transform
}

func (p *TestPattern) SetTransform(transform Matrix) {
	p.Transform = transform
}

func (p *TestPattern) PatternAt(point Tuple) Color {
	return NewColor(point.X, point.Y, point.Z)
}

type StripePattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NewStripePattern(a, b Color) *StripePattern {
	return &StripePattern{a, b, NewIdentity4()}
}

func (p *StripePattern) PatternAt(point Tuple) Color {
	if int(math.Floor(point.X))%2 == 0 {
		return p.A
	}

	return p.B
}

func (p *StripePattern) SetTransform(transform Matrix) {
	p.Transform = transform
}

func (p *StripePattern) GetTransform() Matrix {
	return p.Transform
}

type GradientPattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NewGradientPattern(a, b Color) *GradientPattern {
	return &GradientPattern{a, b, NewIdentity4()}
}

func (p *GradientPattern) PatternAt(point Tuple) Color {
	distance := SubColors(p.B, p.A)
	fraction := point.X - math.Floor(point.X)

	return AddColors(p.A, ColorScalarMultiply(distance, fraction))
}

func (p *GradientPattern) SetTransform(transform Matrix) {
	p.Transform = transform
}

func (p *GradientPattern) GetTransform() Matrix {
	return p.Transform
}

type RingPattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NewRingPattern(a, b Color) *RingPattern {
	return &RingPattern{a, b, NewIdentity4()}
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
}

func (p *RingPattern) GetTransform() Matrix {
	return p.Transform
}

type CheckersPattern struct {
	A         Color
	B         Color
	Transform Matrix
}

func NewCheckersPattern(a, b Color) *CheckersPattern {
	return &CheckersPattern{a, b, NewIdentity4()}
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
}

func (p *CheckersPattern) GetTransform() Matrix {
	return p.Transform
}
