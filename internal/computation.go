package internal

type Computation struct {
	T          float64
	Object     Shape
	Point      Tuple
	OverPoint  Tuple
	UnderPoint Tuple
	EyeV       Tuple
	NormalV    Tuple
	ReflectV   Tuple
	Inside     bool
	N1         float64
	N2         float64
}

func NewComputation() Computation {
	return Computation{}
}

func PrepareComputations(intersection Intersection, ray Ray, xs Intersections) Computation {
	comps := NewComputation()
	comps.T = intersection.T
	comps.Object = intersection.Object
	comps.Point = Position(ray, comps.T)
	comps.EyeV = Negate(ray.Direction)
	comps.NormalV = NormalAt(comps.Object, comps.Point, intersection)
	comps.ReflectV = Reflect(ray.Direction, comps.NormalV)

	if Dot(comps.NormalV, comps.EyeV) < 0 {
		comps.Inside = true
		comps.NormalV = Negate(comps.NormalV)
	} else {
		comps.Inside = false
	}

	comps.OverPoint = AddTuples(comps.Point, TupleScalarMultiply(comps.NormalV, float64EqualityThreshold))
	comps.UnderPoint = SubTuples(comps.Point, TupleScalarMultiply(comps.NormalV, float64EqualityThreshold))

	var containers []Shape

	for _, i := range xs {
		isHit := IntersectEquals(i, intersection)

		if isHit {
			if len(containers) == 0 {
				comps.N1 = 1.0
			} else {
				comps.N1 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}
		}

		index, includes := IndexOf(containers, i.Object)

		if includes {
			containers = DeleteAtIndex(containers, index)
		} else {
			containers = append(containers, i.Object)
		}

		if isHit {
			if len(containers) == 0 {
				comps.N2 = 1.0
			} else {
				comps.N2 = containers[len(containers)-1].GetMaterial().RefractiveIndex
			}

			break
		}
	}

	return comps
}
