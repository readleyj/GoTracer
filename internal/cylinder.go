package internal

import (
	"math"
	"math/rand"
	"time"
)

type Cylinder struct {
	ID               int64
	Material         Material
	Transform        Matrix
	Inverse          Matrix
	InverseTranspose Matrix
	Parent           Shape
	Minimum          float64
	Maximum          float64
	Closed           bool
	HasShadow        bool
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		ID:               rand.Int63(),
		Material:         NewDefaultMaterial(),
		Transform:        NewIdentity4(),
		Inverse:          NewIdentity4(),
		InverseTranspose: NewIdentity4(),
		Parent:           nil,
		Minimum:          math.Inf(-1),
		Maximum:          math.Inf(1),
		Closed:           false,
		HasShadow:        true,
	}
}

func (cyl *Cylinder) GetID() int64 {
	return cyl.ID
}

func (cyl *Cylinder) LocalIntersect(ray Ray) Intersections {
	var xs []Intersection

	a := math.Pow(ray.Direction.X, 2) + math.Pow(ray.Direction.Z, 2)

	if !(math.Abs(a) < float64EqualityThreshold) {
		b := 2*ray.Origin.X*ray.Direction.X + 2*ray.Origin.Z*ray.Direction.Z
		c := math.Pow(ray.Origin.X, 2) + math.Pow(ray.Origin.Z, 2) - 1

		discriminant := b*b - 4*a*c

		if discriminant < 0 {
			return Intersections{}
		}

		t1, t2 := (-b-math.Sqrt(discriminant))/(2*a), (-b+math.Sqrt(discriminant))/(2*a)

		y1 := ray.Origin.Y + t1*ray.Direction.Y
		if cyl.Minimum < y1 && y1 < cyl.Maximum {
			xs = append(xs, NewIntersection(t1, cyl))
		}

		y2 := ray.Origin.Y + t2*ray.Direction.Y
		if cyl.Minimum < y2 && y2 < cyl.Maximum {
			xs = append(xs, NewIntersection(t2, cyl))
		}
	}

	return cyl.IntersectCaps(ray, xs)
}

func (cyl *Cylinder) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	dist := point.X*point.X + point.Z*point.Z

	if dist < 1 && point.Y >= cyl.Maximum-float64EqualityThreshold {
		return NewVector(0, 1, 0)
	} else if dist < 1 && point.Y <= cyl.Minimum+float64EqualityThreshold {
		return NewVector(0, -1, 0)
	}

	return NewVector(point.X, 0, point.Z)
}

func (cyl *Cylinder) GetTransform() Matrix {
	return cyl.Transform
}

func (cyl *Cylinder) SetTransform(transform Matrix) {
	cyl.Transform = transform
	cyl.Inverse = MatrixInverse(cyl.Transform)
	cyl.InverseTranspose = MatrixTranspose(cyl.Inverse)
}

func (cyl *Cylinder) GetInverse() Matrix {
	return cyl.Inverse
}

func (cyl *Cylinder) GetInverseTranspose() Matrix {
	return cyl.InverseTranspose
}

func (cyl *Cylinder) GetMaterial() Material {
	return cyl.Material
}

func (cyl *Cylinder) SetMaterial(material Material) {
	cyl.Material = material
}

func (cyl *Cylinder) IntersectCaps(ray Ray, xs Intersections) Intersections {
	var t float64

	if !cyl.Closed || math.Abs(ray.Direction.Y) < float64EqualityThreshold {
		return xs
	}

	t = (cyl.Minimum - ray.Origin.Y) / ray.Direction.Y
	if CheckCap(ray, t, 1) {
		xs = append(xs, NewIntersection(t, cyl))
	}

	t = (cyl.Maximum - ray.Origin.Y) / ray.Direction.Y
	if CheckCap(ray, t, 1) {
		xs = append(xs, NewIntersection(t, cyl))
	}

	return xs
}

func (cyl *Cylinder) GetParent() Shape {
	return cyl.Parent
}

func (cyl *Cylinder) SetParent(s Shape) {
	cyl.Parent = s
}

func (cyl *Cylinder) CastsShadow() bool {
	return cyl.HasShadow
}
