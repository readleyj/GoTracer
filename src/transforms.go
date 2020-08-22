package internal

import "math"

func Translate(x, y, z float64) Matrix {
	transform := MakeIdentity4()
	transform.Set(x, 0, 3)
	transform.Set(y, 1, 3)
	transform.Set(z, 2, 3)

	return transform
}

func Scale(x, y, z float64) Matrix {
	transform := MakeIdentity4()
	transform.Set(x, 0, 0)
	transform.Set(y, 1, 1)
	transform.Set(z, 2, 2)

	return transform
}

func RotateX(radians float64) Matrix {
	transform := MakeIdentity4()
	transform.Set(math.Cos(radians), 1, 1)
	transform.Set(-math.Sin(radians), 1, 2)
	transform.Set(math.Sin(radians), 2, 1)
	transform.Set(math.Cos(radians), 2, 2)

	return transform
}

func RotateY(radians float64) Matrix {
	transform := MakeIdentity4()
	transform.Set(math.Cos(radians), 0, 0)
	transform.Set(math.Sin(radians), 0, 2)
	transform.Set(-math.Sin(radians), 2, 0)
	transform.Set(math.Cos(radians), 2, 2)

	return transform
}

func RotateZ(radians float64) Matrix {
	transform := MakeIdentity4()
	transform.Set(math.Cos(radians), 0, 0)
	transform.Set(-math.Sin(radians), 0, 1)
	transform.Set(math.Sin(radians), 1, 0)
	transform.Set(math.Cos(radians), 1, 1)

	return transform
}

func Shear(xy, xz, yx, yz, zx, zy float64) Matrix {
	transform := MakeIdentity4()
	transform.Set(xy, 0, 1)
	transform.Set(xz, 0, 2)
	transform.Set(yx, 1, 0)
	transform.Set(yz, 1, 2)
	transform.Set(zx, 2, 0)
	transform.Set(zy, 2, 1)

	return transform
}
