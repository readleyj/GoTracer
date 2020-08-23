package internal

import (
	"math"
	"math/rand"
	"time"
)

type Sphere struct {
	ID        int64
	Transform Matrix
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MakeSphere() *Sphere {
	return &Sphere{rand.Int63(), MakeIdentity4()}
}

func Intersect(s *Sphere, ray Ray) Intersections {
	r := TransformRay(ray, MatrixInverse(s.Transform))
	sphereToRay := SubTuples(r.Origin, Point(0, 0, 0))

	a := Dot(r.Direction, r.Direction)
	b := 2 * Dot(r.Direction, sphereToRay)
	c := Dot(sphereToRay, sphereToRay) - 1
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1, t2 := (-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a)
	i1, i2 := MakeIntersection(t1, s), MakeIntersection(t2, s)

	return MakeIntersections(i1, i2)
}

func (s *Sphere) SetTransform(transform Matrix) {
	s.Transform = transform
}
