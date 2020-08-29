package internal

func NormalAt(s Shape, worldPoint Tuple) Tuple {
	localPoint := WorldToObject(s, worldPoint)
	localNormal := s.LocalNormalAt(localPoint)

	return NormalToWorld(s, localNormal)
}

func Reflect(in, normal Tuple) Tuple {
	return SubTuples(in, TupleScalarMultiply(normal, 2*Dot(in, normal)))
}
