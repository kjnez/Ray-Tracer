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

func (t Tuple) Equals(other Tuple) bool {
	return Equal(t.x, other.x) &&
	Equal(t.y, other.y) &&
	Equal(t.z, other.z) &&
	Equal(t.w, other.w)
}

func (t Tuple) Add(other Tuple) Tuple {
	return Tuple{
		t.x + other.x,
		t.y + other.y,
		t.z + other.z,
		t.w + other.w,
	}
}

func (p Tuple) Subtract(other Tuple) Tuple {
	return Tuple{
		p.x - other.x,
		p.y - other.y,
		p.z - other.z,
		p.w - other.w,
	}
}

func (t Tuple) Negate() Tuple {
	return Tuple{-t.x, -t.y, -t.z, -t.w}
}

