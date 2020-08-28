package internal

import (
	"github.com/google/go-cmp/cmp/cmpopts"
)

const (
	float64EqualityThreshold = 1e-4
	RecursionDepth           = 5
)

func Includes(objects []Shape, target Shape) (int, bool) {
	for index, obj := range objects {
		if ShapesAreIdentical(obj, target) {
			return index, true
		}
	}

	return 0, false
}

func DeleteAtIndex(objects []Shape, index int) []Shape {
	return append(objects[:index], objects[index+1:]...)
}

var opt = cmpopts.EquateApprox(0, float64EqualityThreshold)
