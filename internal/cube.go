package internal

import (
	"math"
	"math/rand"
	"time"
)

type Cube struct {
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

func NewCube() *Cube {
	return &Cube{
		ID:               rand.Int63(),
		Material:         NewDefaultMaterial(),
		Transform:        NewIdentity4(),
		Inverse:          NewIdentity4(),
		InverseTranspose: NewIdentity4(),
		Parent:           nil,
		HasShadow:        true,
	}
}

func (c *Cube) GetID() int64 {
	return c.ID
}

func (c *Cube) LocalIntersect(localRay Ray) Intersections {
	xtMin, xtMax := CheckAxis(localRay.Origin.X, localRay.Direction.X, -1, 1)
	ytMin, ytMax := CheckAxis(localRay.Origin.Y, localRay.Direction.Y, -1, 1)
	ztMin, ztMax := CheckAxis(localRay.Origin.Z, localRay.Direction.Z, -1, 1)

	tMin := math.Max(math.Max(xtMin, ytMin), ztMin)
	tMax := math.Min(math.Min(xtMax, ytMax), ztMax)

	if tMin > tMax {
		return Intersections{}
	}

	i1 := NewIntersection(tMin, c)
	i2 := NewIntersection(tMax, c)

	return NewIntersections(i1, i2)
}

func (c *Cube) LocalNormalAt(point Tuple, hit Intersection) Tuple {
	xAbs, yAbs, zAbs := math.Abs(point.X), math.Abs(point.Y), math.Abs(point.Z)
	maxC := math.Max(math.Max(xAbs, yAbs), zAbs)

	if maxC == xAbs {
		return NewVector(point.X, 0, 0)
	} else if maxC == yAbs {
		return NewVector(0, point.Y, 0)
	}

	return NewVector(0, 0, point.Z)
}

func (c *Cube) GetTransform() Matrix {
	return c.Transform
}

func (c *Cube) SetTransform(transform Matrix) {
	c.Transform = transform
	c.Inverse = MatrixInverse(c.Transform)
	c.InverseTranspose = MatrixTranspose(c.Inverse)
}

func (c *Cube) GetInverse() Matrix {
	return c.Inverse
}

func (c *Cube) GetInverseTranspose() Matrix {
	return c.InverseTranspose
}

func (c *Cube) GetMaterial() Material {
	return c.Material
}

func (c *Cube) SetMaterial(material Material) {
	c.Material = material
}

func (c *Cube) GetParent() Shape {
	return c.Parent
}

func (c *Cube) SetParent(s Shape) {
	c.Parent = s
}

func (c *Cube) CastsShadow() bool {
	return c.HasShadow
}
