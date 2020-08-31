package internal

import (
	"math/rand"
	"sort"
	"time"
)

type CSGOperation int

const (
	CSGUnion CSGOperation = iota
	CSGIntersect
	CSGDifference
)

func (op CSGOperation) String() string {
	return [...]string{"CSGUnion", "CSGIntersection", "CSGDifference"}[op]
}

type CSG struct {
	ID        int64
	Material  Material
	Transform Matrix
	Parent    Shape
	Operation CSGOperation
	Left      Shape
	Right     Shape
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewCSG(op CSGOperation, left, right Shape) *CSG {
	csg := &CSG{
		ID:        rand.Int63(),
		Material:  NewDefaultMaterial(),
		Transform: NewIdentity4(),
		Parent:    nil,
		Operation: op,
		Left:      left,
		Right:     right,
	}

	left.SetParent(csg)
	right.SetParent(csg)

	return csg
}

func (csg *CSG) GetID() int64 {
	return csg.ID
}

func (csg *CSG) LocalIntersect(ray Ray) Intersections {
	leftXS := Intersect(csg.Left, ray)
	rightXS := Intersect(csg.Right, ray)

	xs := append(leftXS, rightXS...)
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].T < xs[j].T
	})

	return FilterIntersections(csg, xs)
}

func (csg *CSG) LocalNormalAt(point Tuple, i Intersection) Tuple {
	return Tuple{}
}

func (csg *CSG) GetTransform() Matrix {
	return csg.Transform
}

func (csg *CSG) SetTransform(transform Matrix) {
	csg.Transform = transform
}

func (csg *CSG) GetMaterial() Material {
	return csg.Material
}

func (csg *CSG) SetMaterial(material Material) {
	csg.Material = material
}

func (csg *CSG) GetParent() Shape {
	return csg.Parent
}

func (csg *CSG) SetParent(s Shape) {
	csg.Parent = s
}

func IntersectionAllowed(op CSGOperation, lhit, inl, inr bool) bool {
	var result bool

	switch op {
	case CSGUnion:
		result = (lhit && !inr) || (!lhit && !inl)

	case CSGIntersect:
		result = (lhit && inr) || (!lhit && inl)

	case CSGDifference:
		result = (lhit && !inr) || (!lhit && inl)
	}

	return result
}

func FilterIntersections(csg *CSG, xs Intersections) Intersections {
	var result Intersections
	inl, inr := false, false

	for _, i := range xs {
		lhit := Includes(csg.Left, i.Object)

		if IntersectionAllowed(csg.Operation, lhit, inl, inr) {
			result = append(result, i)
		}

		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}

	return result
}
