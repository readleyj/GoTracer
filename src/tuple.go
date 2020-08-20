package tuple

import (
	"math"

	"github.com/google/go-cmp/cmp"
)

type Tuple struct {
	X, Y, Z, W float64
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func (t Tuple) IsVector() bool {
	return t.W == 0.0
}

func (t Tuple) IsPoint() bool {
	return t.W == 1.0
}

func TupleEquals(t1, t2 Tuple) bool {
	return cmp.Equal(t1, t2, opt)
}

func (t1 Tuple) Equals(t2 Tuple) bool {
	return cmp.Equal(t1, t2, opt)
}

func Add(t1, t2 Tuple) Tuple {
	return Tuple{t1.X + t2.X, t1.Y + t2.Y, t1.Z + t2.Z, t1.W + t2.W}
}

func (t1 Tuple) Add(t2 Tuple) Tuple {
	return Tuple{t1.X + t2.X, t1.Y + t2.Y, t1.Z + t2.Z, t1.W + t2.W}
}

func Sub(t1, t2 Tuple) Tuple {
	return Tuple{t1.X - t2.X, t1.Y - t2.Y, t1.Z - t2.Z, t1.W - t2.W}
}

func (t1 Tuple) Sub(t2 Tuple) Tuple {
	return Tuple{t1.X - t2.X, t1.Y - t2.Y, t1.Z - t2.Z, t1.W - t2.W}
}

func Negate(t1 Tuple) Tuple {
	return Tuple{-t1.X, -t1.Y, -t1.Z, -t1.W}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.X, -t.Y, -t.Z, -t.W}
}

func MultiplyByScalar(t1 Tuple, scalar float64) Tuple {
	return Tuple{t1.X * scalar, t1.Y * scalar, t1.Z * scalar, t1.W * scalar}
}

func (t Tuple) MultiplyByScalar(scalar float64) Tuple {
	return Tuple{t.X * scalar, t.Y * scalar, t.Z * scalar, t.W * scalar}
}

func DivideByScalar(t1 Tuple, scalar float64) Tuple {
	return Tuple{t1.X / scalar, t1.Y / scalar, t1.Z / scalar, t1.W / scalar}
}

func (t Tuple) DivideByScalar(scalar float64) Tuple {
	return Tuple{t.X / scalar, t.Y / scalar, t.Z / scalar, t.W / scalar}
}

func Magnitude(t1 Tuple) float64 {
	return math.Sqrt(t1.X*t1.X + t1.Y*t1.Y + t1.Z*t1.Z)
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z)
}

func Normalize(t1 Tuple) Tuple {
	magnitude := Magnitude(t1)
	return Tuple{t1.X / magnitude, t1.Y / magnitude, t1.Z / magnitude, t1.W / magnitude}
}

func (t Tuple) Normalize() Tuple {
	magnitude := Magnitude(t)
	return Tuple{t.X / magnitude, t.Y / magnitude, t.Z / magnitude, t.W / magnitude}
}

func Dot(t1, t2 Tuple) float64 {
	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z + t1.W*t2.W
}

func (t1 Tuple) Dot(t2 Tuple) float64 {
	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z + t1.W*t2.W
}

func Cross(t1, t2 Tuple) Tuple {
	return Vector(
		t1.Y*t2.Z-t1.Z*t2.Y,
		t1.Z*t2.X-t1.X*t2.Z,
		t1.X*t2.Y-t1.Y*t2.X,
	)
}

func (t1 Tuple) Cross(t2 Tuple) Tuple {
	return Vector(
		t1.Y*t2.Z-t1.Z*t2.Y,
		t1.Z*t2.X-t1.X*t2.Z,
		t1.X*t2.Y-t1.Y*t2.X,
	)
}