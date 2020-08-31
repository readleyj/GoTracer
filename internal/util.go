package internal

import (
	"math"

	"github.com/google/go-cmp/cmp/cmpopts"
)

const (
	float64EqualityThreshold = 1e-4
	RecursionDepth           = 5
)

var opt = cmpopts.EquateApprox(0, float64EqualityThreshold)

func Includes(left Shape, target Shape) bool {
	switch t := left.(type) {
	case *Group:
		for _, child := range t.Children {
			if child.GetID() == target.GetID() {
				return true
			}

			return Includes(child, target)
		}

	case *CSG:
		includesLeft := Includes(t.Left, target)
		includesRight := Includes(t.Right, target)
		return includesLeft || includesRight

	default:
		return left.GetID() == target.GetID()
	}

	return false
}

func IndexOf(objects []Shape, target Shape) (int, bool) {
	for index, obj := range objects {
		if obj.GetID() == target.GetID() {
			return index, true
		}
	}

	return 0, false
}

func DeleteAtIndex(objects []Shape, index int) []Shape {
	return append(objects[:index], objects[index+1:]...)
}

func CheckAxis(origin, direction float64) (float64, float64) {
	var tMin, tMax float64

	tMinNumerator := -origin - 1
	tMaxNumerator := -origin + 1

	if math.Abs(direction) >= float64EqualityThreshold {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	} else {
		tMin = tMinNumerator * math.Inf(1)
		tMax = tMaxNumerator * math.Inf(1)
	}

	if tMin > tMax {
		tMin, tMax = tMax, tMin
	}

	return tMin, tMax
}

func CheckCap(r Ray, t, boundary float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (x*x + z*z) <= math.Abs(boundary)
}
