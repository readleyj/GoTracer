package internal

type Ray struct {
	origin, direction Tuple
}

func Position(r Ray, dist float64) Tuple {
	return AddTuples(r.origin, r.direction.MultiplyByScalar(dist))
}

func TransformRay(r Ray, transform Matrix) Ray {
	newOrigin := MatrixTupleMultiply(transform, r.origin)
	newDirection := MatrixTupleMultiply(transform, r.direction)

	return Ray{newOrigin, newDirection}
}
