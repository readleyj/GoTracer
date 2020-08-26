package internal

import (
	"math"
	"math/rand"
	"time"

	"github.com/google/go-cmp/cmp"
)

type Sphere struct {
	ID        int64
	Material  Material
	Transform Matrix
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewSphere() *Sphere {
	return &Sphere{
		rand.Int63(),
		NewDefaultMaterial(),
		NewIdentity4(),
	}
}

func Intersect(s *Sphere, ray Ray) Intersections {
	r := TransformRay(ray, MatrixInverse(s.Transform))
	sphereToRay := SubTuples(r.Origin, NewPoint(0, 0, 0))

	a := Dot(r.Direction, r.Direction)
	b := 2 * Dot(r.Direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1, t2 := (-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a)
	i1, i2 := NewIntersection(t1, s), NewIntersection(t2, s)

	return NewIntersections(i1, i2)
}

func SphereEquals(s1, s2 *Sphere) bool {
	return cmp.Equal(s1.Material, s2.Material, opt) && MatrixEquals(s1.Transform, s2.Transform)
}

func (s *Sphere) SetTransform(transform Matrix) {
	s.Transform = transform
}

func (s *Sphere) SetMaterial(material Material) {
	s.Material = material
}
