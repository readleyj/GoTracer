package internal

import (
	"math"
	"math/rand"
	"time"
)

type Cube struct {
	ID        int64
	Material  Material
	Transform Matrix
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewCube() *Cube {
	return &Cube{
		rand.Int63(),
		NewDefaultMaterial(),
		NewIdentity4(),
	}
}

func (c *Cube) GetID() int64 {
	return c.ID
}

func (c *Cube) LocalIntersect(localRay Ray) Intersections {
	xtMin, xtMax := CheckAxis(localRay.Origin.X, localRay.Direction.X)
	ytMin, ytMax := CheckAxis(localRay.Origin.Y, localRay.Direction.Y)
	ztMin, ztMax := CheckAxis(localRay.Origin.Z, localRay.Direction.Z)

	tMin := math.Max(math.Max(xtMin, ytMin), ztMin)
	tMax := math.Min(math.Min(xtMax, ytMax), ztMax)

	if tMin > tMax {
		return Intersections{}
	}

	i1 := NewIntersection(tMin, c)
	i2 := NewIntersection(tMax, c)

	return NewIntersections(i1, i2)
}

func (c *Cube) LocalNormalAt(point Tuple) Tuple {
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
}

func (c *Cube) GetMaterial() Material {
	return c.Material
}

func (c *Cube) SetMaterial(material Material) {
	c.Material = material
}
