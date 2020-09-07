package internal

import "math"

type BoundingBox struct {
	Min Tuple
	Max Tuple
}

func NewEmptyBoundingBox() BoundingBox {
	return BoundingBox{
		Min: NewPoint(math.Inf(1), math.Inf(1), math.Inf(1)),
		Max: NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1)),
	}
}

func NewBoundingBox(min, max Tuple) BoundingBox {
	return BoundingBox{
		Min: min,
		Max: max,
	}
}

func (box *BoundingBox) AddPoint(point Tuple) {
	box.Min.X = math.Min(box.Min.X, point.X)
	box.Min.Y = math.Min(box.Min.Y, point.Y)
	box.Min.Z = math.Min(box.Min.Z, point.Z)

	box.Max.X = math.Max(box.Max.X, point.X)
	box.Max.Y = math.Max(box.Max.Y, point.Y)
	box.Max.Z = math.Max(box.Max.Z, point.Z)
}

func (box1 *BoundingBox) AddBox(box2 BoundingBox) {
	box1.AddPoint(box2.Min)
	box1.AddPoint(box2.Max)
}

func (box *BoundingBox) ContainsPoint(point Tuple) bool {
	return (point.X >= box.Min.X && point.X <= box.Max.X) &&
		(point.Y >= box.Min.Y && point.Y <= box.Max.Y) &&
		(point.Z >= box.Min.Z && point.Z <= box.Max.Z)
}

func (box1 *BoundingBox) ContainsBox(box2 BoundingBox) bool {
	return box1.ContainsPoint(box2.Min) && box1.ContainsPoint(box2.Max)
}

func BoundsOf(shape Shape) BoundingBox {
	switch temp := shape.(type) {
	case *Sphere, *TestShape:
		return NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
	case *Plane:
		return NewBoundingBox(NewPoint(math.Inf(-1), 0, math.Inf(-1)), NewPoint(math.Inf(1), 0, math.Inf(1)))
	case *Cube:
		return NewBoundingBox(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
	case *Cylinder:
		return NewBoundingBox(NewPoint(-1, temp.Minimum, -1), NewPoint(1, temp.Maximum, 1))
	case *Cone:
		a := math.Abs(temp.Minimum)
		b := math.Abs(temp.Maximum)
		limit := math.Max(a, b)
		return NewBoundingBox(NewPoint(-limit, temp.Minimum, -limit), NewPoint(limit, temp.Maximum, limit))
	case *Triangle:
		box := NewEmptyBoundingBox()
		box.AddPoint(temp.P1)
		box.AddPoint(temp.P2)
		box.AddPoint(temp.P3)

		return box
	case *Group:
		box := NewEmptyBoundingBox()

		for _, child := range temp.Children {
			cBox := ParentSpaceBoundsOf(child)
			box.AddBox(cBox)
		}

		return box
	case *CSG:
		box := NewEmptyBoundingBox()
		box.AddBox(ParentSpaceBoundsOf(temp.Left))
		box.AddBox(ParentSpaceBoundsOf(temp.Right))
		return box
	default:
		panic("Invalid shape")
	}
}

func TransformBox(box BoundingBox, transform Matrix) BoundingBox {
	p1 := box.Min
	p2 := NewPoint(box.Min.X, box.Min.Y, box.Max.Z)
	p3 := NewPoint(box.Min.X, box.Max.Y, box.Min.Z)
	p4 := NewPoint(box.Min.X, box.Max.Y, box.Max.Z)
	p5 := NewPoint(box.Max.X, box.Min.Y, box.Min.Z)
	p6 := NewPoint(box.Max.X, box.Min.Y, box.Max.Z)
	p7 := NewPoint(box.Max.X, box.Max.Y, box.Min.Z)
	p8 := box.Max
	points := []Tuple{p1, p2, p3, p4, p5, p6, p7, p8}

	newBox := NewEmptyBoundingBox()

	for _, point := range points {
		newBox.AddPoint(MatrixTupleMultiply(transform, point))
	}

	return newBox
}

func ParentSpaceBoundsOf(shape Shape) BoundingBox {
	return TransformBox(BoundsOf(shape), shape.GetTransform())
}
