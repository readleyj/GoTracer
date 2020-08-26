package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrecompIntersectionState(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(4, shape)
	comps := PrepareComputations(i, r)

	assert.InDelta(t, i.T, comps.T, float64EqualityThreshold)
	assert.True(t, SphereEquals(i.Object, comps.Object))
	assert.True(t, TupleEquals(NewPoint(0, 0, -1), comps.Point))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), comps.EyeV))
	assert.True(t, TupleEquals(NewVector(0, 0, -1), comps.NormalV))
}
