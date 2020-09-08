package internal

import "math"

func SplitBounds(box BoundingBox) (BoundingBox, BoundingBox) {
	dx := box.Max.X - box.Min.X
	dy := box.Max.Y - box.Min.Y
	dz := box.Max.Z - box.Min.Z

	greatest := math.Max(math.Max(dx, dy), dz)

	x0, y0, z0 := box.Min.X, box.Min.Y, box.Min.Z
	x1, y1, z1 := box.Max.X, box.Max.Y, box.Max.Z

	if greatest == dx {
		x0 += dx / 2.0
		x1 = x0
	} else if greatest == dy {
		y0 += dy / 2.0
		y1 = y0
	} else {
		z0 += dz / 2.0
		z1 = z0
	}

	midMin := NewPoint(x0, y0, z0)
	midMax := NewPoint(x1, y1, z1)

	left := NewBoundingBox(box.Min, midMax)
	right := NewBoundingBox(midMin, box.Max)

	return left, right
}

func PartitionChildren(group *Group) (*Group, *Group) {
	var rest []Shape

	left := NewGroup()
	right := NewGroup()
	bounds := BoundsOf(group)
	leftBounds, rightBounds := SplitBounds(bounds)

	for _, child := range group.Children {
		childBound := ParentSpaceBoundsOf(child)

		if leftBounds.ContainsBox(childBound) {
			left.AddChild(child)
		} else if rightBounds.ContainsBox(childBound) {
			right.AddChild(child)
		} else {
			rest = append(rest, child)
		}
	}

	group.Children = group.Children[:0]
	group.Children = append(group.Children, rest...)

	return left, right
}

func MakeSubgroup(group *Group, shapes ...Shape) {
	subgroup := NewGroup()

	for _, shape := range shapes {
		subgroup.AddChild(shape)
	}

	group.AddChild(subgroup)
}

func Divide(shape Shape, threshold int) {
	switch g := shape.(type) {
	case *Group:
		if threshold <= len(g.Children) {
			left, right := PartitionChildren(g)

			if len(left.Children) != 0 {
				MakeSubgroup(g, left.Children...)
			}
			if len(right.Children) != 0 {
				MakeSubgroup(g, right.Children...)
			}
		}

		for _, child := range g.Children {
			Divide(child, threshold)
		}
	case *CSG:
		Divide(g.Left, threshold)
		Divide(g.Right, threshold)
	default:

	}
}
