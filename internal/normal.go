package internal

func NormalAt(s *Sphere, worldPoint Tuple) Tuple {
	inverseTransform := MatrixInverse(s.Transform)
	objectPoint := MatrixTupleMultiply(inverseTransform, worldPoint)
	objectNormal := SubTuples(objectPoint, NewPoint(0, 0, 0))
	worldNormal := MatrixTupleMultiply(MatrixTranspose(inverseTransform), objectNormal)
	worldNormal.W = 0

	return Normalize(worldNormal)
}

func Reflect(in, normal Tuple) Tuple {
	return SubTuples(in, TupleScalarMultiply(normal, 2*Dot(in, normal)))
}
