package internal

type Computation struct {
	T         float64
	Object    *Sphere
	Point     Tuple
	OverPoint Tuple
	EyeV      Tuple
	NormalV   Tuple
	Inside    bool
}

func NewComputation() Computation {
	return Computation{}
}

func PrepareComputations(intersection Intersection, ray Ray) Computation {
	comps := NewComputation()
	comps.T = intersection.T
	comps.Object = intersection.Object
	comps.Point = Position(ray, comps.T)
	comps.EyeV = Negate(ray.Direction)
	comps.NormalV = NormalAt(comps.Object, comps.Point)

	if Dot(comps.NormalV, comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = Negate(comps.NormalV)
	} else {
		comps.Inside = false
	}

	comps.OverPoint = AddTuples(comps.Point, TupleScalarMultiply(comps.NormalV, float64EqualityThreshold))

	return comps
}
