package internal

import (
	"math"
	"math/rand"
	"time"
)

type Cone struct {
	ID        int64
	Material  Material
	Transform Matrix
	Parent    Shape
	Minimum   float64
	Maximum   float64
	Closed    bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewCone() *Cone {
	return &Cone{
		ID:        rand.Int63(),
		Material:  NewDefaultMaterial(),
		Transform: NewIdentity4(),
		Parent:    nil,
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
	}
}

func (cone *Cone) GetID() int64 {
	return cone.ID
}

func (cone *Cone) LocalIntersect(ray Ray) Intersections {
	var xs []Intersection

	direction, origin := ray.Direction, ray.Origin
	a := math.Pow(direction.X, 2) - math.Pow(direction.Y, 2) + math.Pow(direction.Z, 2)
	b := 2*origin.X*direction.X - 2*origin.Y*direction.Y + 2*origin.Z*direction.Z
	c := math.Pow(origin.X, 2) - math.Pow(origin.Y, 2) + math.Pow(origin.Z, 2)
	discriminant := b*b - 4*a*c

	if (math.Abs(a) < float64EqualityThreshold && math.Abs(b) < float64EqualityThreshold) || discriminant < 0 {
		return xs
	}

	var t1, t2 float64

	if math.Abs(a) < float64EqualityThreshold {
		t1 = -c / (2 * b)
		y1 := ray.Origin.Y + t1*ray.Direction.Y

		if y1 > cone.Minimum && y1 < cone.Maximum {
			xs = append(xs, NewIntersection(t1, cone))
		}

	} else {
		t1, t2 = (-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a)

		y1 := ray.Origin.Y + t1*ray.Direction.Y
		if cone.Minimum < y1 && y1 < cone.Maximum {
			xs = append(xs, NewIntersection(t1, cone))
		}

		y2 := ray.Origin.Y + t2*ray.Direction.Y
		if cone.Minimum < y2 && y2 < cone.Maximum {
			xs = append(xs, NewIntersection(t2, cone))
		}
	}

	return cone.IntersectCaps(ray, xs)
}

func (cone *Cone) LocalNormalAt(point Tuple, i Intersection) Tuple {
	dist := point.X*point.X + point.Z*point.Z

	if dist < 1 && point.Y >= cone.Maximum-float64EqualityThreshold {
		return NewVector(0, 1, 0)
	} else if dist < 1 && point.Y <= cone.Minimum+float64EqualityThreshold {
		return NewVector(0, -1, 0)
	}

	y := math.Sqrt(dist)

	if point.Y > 0 {
		y = -y
	}

	return NewVector(point.X, y, point.Z)
}

func (cone *Cone) GetTransform() Matrix {
	return cone.Transform
}

func (cone *Cone) SetTransform(transform Matrix) {
	cone.Transform = transform
}

func (cone *Cone) GetMaterial() Material {
	return cone.Material
}

func (cone *Cone) SetMaterial(material Material) {
	cone.Material = material
}

func (cone *Cone) IntersectCaps(ray Ray, xs Intersections) Intersections {
	var t float64

	if !cone.Closed || math.Abs(ray.Direction.Y) < float64EqualityThreshold {
		return xs
	}

	t = (cone.Minimum - ray.Origin.Y) / ray.Direction.Y
	if CheckCap(ray, t, cone.Minimum) {
		xs = append(xs, NewIntersection(t, cone))
	}

	t = (cone.Maximum - ray.Origin.Y) / ray.Direction.Y
	if CheckCap(ray, t, cone.Maximum) {
		xs = append(xs, NewIntersection(t, cone))
	}

	return xs
}

func (cone *Cone) GetParent() Shape {
	return cone.Parent
}

func (cone *Cone) SetParent(s Shape) {
	cone.Parent = s
}
