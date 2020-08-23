package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSphereDefaultTransform(t *testing.T) {
	s := MakeSphere()
	assert.True(t, MatrixEquals(s.Transform, Identity4))
}

func TestSphereSetTransform(t *testing.T) {
	s := MakeSphere()
	transform := Translate(2, 3, 4)
	s.SetTransform(transform)
	assert.True(t, MatrixEquals(s.Transform, transform))
}
