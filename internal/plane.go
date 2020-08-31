package internal

import (
	"math"
	"math/rand"
	"time"
)

type Plane struct {
	ID        int64
	Material  Material
	Transform Matrix
	Parent    Shape
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewPlane() *Plane {
	return &Plane{
		rand.Int63(),
		NewDefaultMaterial(),
		NewIdentity4(),
		nil,
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
