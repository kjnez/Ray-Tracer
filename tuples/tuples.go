package tuples

import (
	"math"
)

type Tuple struct {
	x, y, z, w float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func (t Tuple) IsPoint() bool {
	return t.w == 1.0
}

func (t Tuple) IsVector() bool {
	return t.w == 0.0
}

const EPSILON = 1e-5

func Equal(a, b float64) bool {
	if math.Abs(a - b) < EPSILON {
		return true
	} else {
		return false
	}
}

func Equals(t1, t2 Tuple) bool {
	return Equal(t1.x, t2.x) &&
	Equal(t1.y, t2.y) &&
	Equal(t1.z, t2.z) &&
	Equal(t1.w, t2.w)
}


func Add(t1, t2 Tuple) Tuple {
	return Tuple{
		t1.x + t2.x,
		t1.y + t2.y,
		t1.z + t2.z,
		t1.w + t2.w,
	}
}

func Subtract(t1, t2 Tuple) Tuple {
	return Tuple{
		t1.x - t2.x,
		t1.y - t2.y,
		t1.z - t2.z,
		t1.w - t2.w,
	}
}

func Negate(t Tuple) Tuple {
	return Tuple{-t.x, -t.y, -t.z, -t.w}
}

func Multiply(tuple Tuple, scalar float64) Tuple {
	return Tuple{
		tuple.x * scalar,
		tuple.y * scalar,
		tuple.z * scalar,
		tuple.w * scalar,
	}
}

func Divide(tuple Tuple, scalar float64) Tuple {
	if scalar == 0 {
		panic("division by zero")
	}
	return Multiply(tuple, 1.0 / scalar)
}
