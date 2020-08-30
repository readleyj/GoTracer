package internal

import (
	"math"
	"math/rand"
	"time"
)

type Sphere struct {
	ID        int64
	Material  Material
	Transform Matrix
	Parent    *Group
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewSphere() *Sphere {
	return &Sphere{
		rand.Int63(),
		NewDefaultMaterial(),
		NewIdentity4(),
		nil,
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
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}

func (s *Sphere) SetMaterial(material Material) {
	s.Material = material
}

func (s *Sphere) GetParent() *Group {
	return s.Parent
}

func (s *Sphere) SetParent(g *Group) {
	s.Parent = g
}
