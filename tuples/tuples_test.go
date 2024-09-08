package tuples

import (
	"testing"
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
	if !p.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, p)
	}
}

func TestVectorCreatesW0Tuple(t *testing.T) {
	v := Vector(4, -4, 3)
	expected := NewTuple(4, -4, 3, 0)
	if !v.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, v)
	}
}

func TestAddTwoTuples(t *testing.T) {
	t1 := NewTuple(3, -2, 5, 1)
	t2 := NewTuple(-2, 3, 1, 0)
	sum := t1.Add(t2)
	expected := NewTuple(1, 1, 6, 1)
	if !sum.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, sum)
	}
}

func TestSubtractTwoPoints(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)
	subtracted := p1.Subtract(p2)
	expected := Vector(-2, -4, -6)
	if !subtracted.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, subtracted)
	}
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p := Point(3, 2, 1)
	v := Vector(5, 6, 7)
	subtracted := p.Subtract(v)
	expected := Point(-2, -4, -6)
	if !subtracted.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, subtracted)
	}
}

func TestSutractVectorFromVector(t *testing.T) {
	v1 := Vector(3, 2, 1)
	v2 := Vector(5, 6, 7)
	subtracted := v1.Subtract(v2)
	expected := Vector(-2, -4, -6)
	if !subtracted.Equals(expected) {
		t.Errorf("Expected %v, got %v", expected, subtracted)
	}
}
