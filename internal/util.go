package internal

import (
	"github.com/google/go-cmp/cmp/cmpopts"
)

const float64EqualityThreshold = 1e-4

var opt = cmpopts.EquateApprox(0, float64EqualityThreshold)
