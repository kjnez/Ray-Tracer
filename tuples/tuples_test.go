package tuples

import (
	"testing"
	"math"
)

func TestTupleWithW1IsPoint(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	if a.x != 4.3 || a.y != -4.2 || a.z != 3.1 {
		t.Errorf("Expected (4.3, -4.2, 3.1), got (%f, %f, %f)", a.x, a.y, a.z)
	}
	if !a.IsPoint() {
		t.Error("Expected a to be a point")
	}
	if a.IsVector() {
		t.Error("Expected a not to be a vector")
	}
}

func TestTupleWithW0IsVector(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 0.0)
	if a.x != 4.3 || a.y != -4.2 || a.z != 3.1 || a.w != 0.0 {
		t.Errorf("Expected (4.3, -4.2, 3.1, 0.0), got (%f, %f, %f, %f)", a.x, a.y, a.z, a.w)
	}
	if a.IsPoint() {
		t.Error("Expected a not to be a point")
	}
	if !a.IsVector() {
		t.Error("Expected a to be a vector")
	}
}

func TestPointCreatesW1Tuple(t *testing.T) {
	p := Point(4, -4, 3)
	expected := NewTuple(4, -4, 3, 1)
	if !Equals(p, expected) {
		t.Errorf("Expected %v, got %v", expected, p)
	}
}

func TestVectorCreatesW0Tuple(t *testing.T) {
	v := Vector(4, -4, 3)
	expected := NewTuple(4, -4, 3, 0)
	if !Equals(v, expected) {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestAddTwoTuples(t *testing.T) {
	t1 := NewTuple(3, -2, 5, 1)
	t2 := NewTuple(-2, 3, 1, 0)
	sum := Add(t1, t2)
	expected := NewTuple(1, 1, 6, 1)
	if !Equals(sum, expected) {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}

func TestSubtractTwoPoints(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	subtracted := Subtract(p1, p2)
	expected := Vector(-2, -4, -6)
	if !Equals(subtracted, expected) {
		t.Errorf("Expected %v, got %v", expected, subtracted)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	subtracted := Subtract(p, v)
	expected := Point(-2, -4, -6)
	if !Equals(subtracted, expected) {
		t.Errorf("Expected %v, got %v", expected, subtracted)
	}
}

func TestSutractVectorFromVector(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	subtracted := Subtract(v1, v2)
	expected := Vector(-2, -4, -6)
	if !Equals(subtracted, expected) {
		t.Errorf("Expected %v, got %v", expected, subtracted)
	}
}

func TestNegateTuple(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)
	negateive_t := Negate(tuple)
	expected := NewTuple(-1, 2, -3, 4)
	if !Equals(negateive_t, expected) {
		t.Errorf("Expected %v, got %v", expected, negateive_t)
	}
}

func TestMultiply(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)
	scalar1 := 3.5
	scalar2 := 0.5
	multiply1 := Multiply(tuple, scalar1)
	multiply2 := Multiply(tuple, scalar2)
	expected1 := Tuple{3.5, -7, 10.5, -14}
	expected2 := Tuple{0.5, -1, 1.5, -2}
	if !Equals(multiply1, expected1) {
		t.Errorf("Expected %v, got %v", expected1, multiply1)
	}
	if !Equals(multiply2, expected2) {
		t.Errorf("Expected %v, got %v", expected2, multiply2)
	}
}

func TestDivide(t *testing.T) {
	tuple := NewTuple(1, -2, 3, -4)
	scalar := 2.0
	division := Divide(tuple, scalar)
	expected := Tuple{0.5, -1, 1.5, -2}
	if !Equals(division, expected) {
		t.Errorf("Expected %v, got %v", expected, division)
	}
}

func TestMagnitude(t *testing.T) {
	tests := []struct {
		name     string
		vector   Tuple
		expected float64
	}{
		{"Unit X", Vector(1, 0, 0), 1.0},
		{"Unit Y", Vector(0, 1, 0), 1.0},
		{"Unit Z", Vector(0, 0, 1), 1.0},
		{"Vector 1,2,3", Vector(1, 2, 3), math.Sqrt(14)},
		{"Vector -1,-2,-3", Vector(-1, -2, -3), math.Sqrt(14)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Magnitude(tt.vector)
			if !Equal(got, tt.expected) {
				t.Errorf("Magnitude() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	v1, v2 := Vector(4, 0, 0), Vector(1, 2, 3)
	nv1, nv2 := Normalize(v1), Normalize(v2)
	ev1, ev2 := Vector(1, 0, 0), Vector(1 / math.Sqrt(14), 2 / math.Sqrt(14), 3 / math.Sqrt(14))
	m2, em2 := Magnitude(nv2), 1.0
	if !Equals(nv1, ev1) {
		t.Errorf("Expected %v, got %v", ev1, nv1)
	}
	if !Equals(nv2, ev2) {
		t.Errorf("Expected %v, got %v", ev2, nv2)
	}
	if !Equal(m2, em2) {
		t.Errorf("Expected %v, got %v", em2, m2)
	}
}

func TestDotProduct(t *testing.T) {
	v1, v2 := Vector(1, 2, 3), Vector(2, 3, 4)
	dotProduct := DotProduct(v1, v2)
	expected := 20.0
	if !Equal(dotProduct, expected) {
		t.Errorf("Expected %v, got %v", expected, dotProduct)
	}
}

func TestCrossProduct(t *testing.T) {
	v1, v2 := Vector(1, 2, 3), Vector(2, 3, 4)
	crossProductV1V2 := CrossProduct(v1, v2)
	crossProductV2V1 := CrossProduct(v2, v1)
	expectedV1V2 := Vector(-1, 2, -1)
	expectedV2V1 := Vector(1, -2, 1)
	if !Equals(crossProductV1V2, expectedV1V2) {
		t.Errorf("Expected %v, got %v", expectedV1V2, crossProductV1V2)
	}

	if !Equals(crossProductV2V1, expectedV2V1) {
		t.Errorf("Expected %v, got %v", expectedV2V1, crossProductV2V1)
	}
}

func TestColor(t *testing.T) {
	c := Color{-0.5, 0.4, 1.7}
	if c.red != -0.5 {
		t.Errorf("Expected %v, got %v", -0.5, c.red)
	}
	if c.green != 0.4 {
		t.Errorf("Expected %v, got %v", 0.4, c.green)
	}
	if c.blue != 1.7 {
		t.Errorf("Expected %v, got %v", 1.7, c.blue)
	}
}

func TestAddColor(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	sum := AddColor(c1, c2)
	expected := Color{1.6, 0.7, 1.0}
	if !EqualsColor(sum, expected) {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}

func TestSubtractColor(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	diff := SubtractColor(c1, c2)
	expected := Color{0.2, 0.5, 0.5}
	if !EqualsColor(expected, diff) {
		t.Errorf("Expected %v, got %v", expected, diff)
	}
}
