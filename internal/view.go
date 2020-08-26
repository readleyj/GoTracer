package internal

func ViewTransform(from, to, up Tuple) Matrix {
	forward := Normalize(SubTuples(to, from))
	left := Cross(forward, Normalize(up))
	trueUp := Cross(left, forward)

	orientation := NewIdentity4()

	orientation.Set(left.X, 0, 0)
	orientation.Set(left.Y, 0, 1)
	orientation.Set(left.Z, 0, 2)

	orientation.Set(trueUp.X, 1, 0)
	orientation.Set(trueUp.Y, 1, 1)
	orientation.Set(trueUp.Z, 1, 2)

	orientation.Set(-forward.X, 2, 0)
	orientation.Set(-forward.Y, 2, 1)
	orientation.Set(-forward.Z, 2, 2)

	return MatrixMultiply(orientation, Translate(-from.X, -from.Y, -from.Z))
}
