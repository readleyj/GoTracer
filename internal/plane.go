package internal

import (
	"math"
	"math/rand"
	"time"
)

type Plane struct {
	ID               int64
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
	Parent           Shape
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewPlane() *Plane {
	return &Plane{
		ID:               rand.Int63(),
		Material:         NewDefaultMaterial(),
		Transform:        NewIdentity4(),
		Inverse:          NewIdentity4(),
		InverseTranspose: NewIdentity4(),
		Parent:           nil,
	}
}

func (p *Plane) GetID() int64 {
	return p.ID
}

func (p *Plane) LocalIntersect(localRay Ray) Intersections {
	if math.Abs(localRay.Direction.Y) < float64EqualityThreshold {
		return Intersections{}
	}

	t := -localRay.Origin.Y / localRay.Direction.Y
	return NewIntersections(NewIntersection(t, p))
}

func (p *Plane) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return NewVector(0, 1, 0)
}

func (p *Plane) GetTransform() Matrix {
	return p.Transform
}

func (p *Plane) SetTransform(transform Matrix) {
	p.Transform = transform
	p.Inverse = MatrixInverse(p.Transform)
	p.InverseTranspose = MatrixTranspose(p.Inverse)
}

func (p *Plane) GetInverse() Matrix {
	return p.Inverse
}

func (p *Plane) GetInverseTranspose() Matrix {
	return p.InverseTranspose
}

func (p *Plane) GetMaterial() Material {
	return p.Material
}

func (p *Plane) SetMaterial(material Material) {
	p.Material = material
}

func (p *Plane) GetParent() Shape {
	return p.Parent
}

func (p *Plane) SetParent(s Shape) {
	p.Parent = s
}
