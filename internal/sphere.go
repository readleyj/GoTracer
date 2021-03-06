package internal

import (
	"math"
	"math/rand"
	"time"
)

type Sphere struct {
	ID               int64
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
	Parent           Shape
	HasShadow        bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewSphere() *Sphere {
	return &Sphere{
		ID:               rand.Int63(),
		Material:         NewDefaultMaterial(),
		Transform:        NewIdentity4(),
		Inverse:          NewIdentity4(),
		InverseTranspose: NewIdentity4(),
		Parent:           nil,
		HasShadow:        true,
	}
}

func NewGlassSphere() *Sphere {
	s := NewSphere()
	s.Material.Transparency = 1.0
	s.Material.RefractiveIndex = 1.5

	return s
}

func (s *Sphere) GetID() int64 {
	return s.ID
}

func (s *Sphere) LocalIntersect(localRay Ray) Intersections {
	sphereToRay := SubTuples(localRay.Origin, NewPoint(0, 0, 0))
	a := Dot(localRay.Direction, localRay.Direction)
	b := 2 * Dot(localRay.Direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1, t2 := (-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a)
	i1, i2 := NewIntersection(t1, s), NewIntersection(t2, s)

	return NewIntersections(i1, i2)
}

func (s *Sphere) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	return SubTuples(point, NewPoint(0, 0, 0))
}

func (s *Sphere) GetTransform() Matrix {
	return s.Transform
}

func (s *Sphere) SetTransform(transform Matrix) {
	s.Transform = transform
	s.Inverse = MatrixInverse(s.Transform)
	s.InverseTranspose = MatrixTranspose(s.Inverse)
}

func (s *Sphere) GetInverse() Matrix {
	return s.Inverse
}

func (s *Sphere) GetInverseTranspose() Matrix {
	return s.InverseTranspose
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) SetMaterial(material Material) {
	s.Material = material
}

func (s *Sphere) GetParent() Shape {
	return s.Parent
}

func (s *Sphere) SetParent(shape Shape) {
	s.Parent = shape
}

func (s *Sphere) CastsShadow() bool {
	return s.HasShadow
}
