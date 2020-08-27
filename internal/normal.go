package internal

func NormalAt(s Shape, worldPoint Tuple) Tuple {
	inverseTransform := MatrixInverse(s.GetTransform())
	localPoint := MatrixTupleMultiply(inverseTransform, worldPoint)
	localNormal := s.LocalNormalAt(localPoint)
	worldNormal := MatrixTupleMultiply(MatrixTranspose(inverseTransform), localNormal)
	worldNormal.W = 0

	return Normalize(worldNormal)
}

func Reflect(in, normal Tuple) Tuple {
	return SubTuples(in, TupleScalarMultiply(normal, 2*Dot(in, normal)))
}
