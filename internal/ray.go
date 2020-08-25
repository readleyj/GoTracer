package internal

type Ray struct {
	Origin, Direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func Position(r Ray, dist float64) Tuple {
	return AddTuples(r.Origin, r.Direction.MultiplyByScalar(dist))
}

func TransformRay(r Ray, transform Matrix) Ray {
	newOrigin := MatrixTupleMultiply(transform, r.Origin)
	newDirection := MatrixTupleMultiply(transform, r.Direction)

	return NewRay(newOrigin, newDirection)
}
