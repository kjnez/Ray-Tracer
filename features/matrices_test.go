package features

import (
	"fmt"
	"testing"
)

func Test4x4Matrix(t *testing.T) {
	mat := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	testCases := []struct {
		row int
		col int
		expected float64
	}{
		{0, 0, 1},
		{0, 3, 4},
		{1, 0, 5.5},
		{1, 2, 7.5},
		{2, 2, 11},
		{3, 0, 13.5},
		{3, 2, 15.5},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("row=%d,col=%d", tc.row, tc.col), func(t *testing.T) {
			if mat[tc.row][tc.col] != tc.expected {
				t.Errorf("Expected %f, got %f", tc.expected, mat[tc.row][tc.col])
			}
		})
	}
}

func Test2x2Matrix(t *testing.T) {
	mat := [][]float64{
		{-3, 5},
		{1, -2},
	}
	testCases := []struct {
		row int
		col int
		expected float64
	}{
		{0, 0, -3},
		{0, 1, 5},
		{1, 0, 1},
		{1, 1, -2},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("row=%d,col=%d", tc.row, tc.col), func(t *testing.T) {
			if mat[tc.row][tc.col] != tc.expected {
				t.Errorf("Expected %f, got %f", tc.expected, mat[tc.row][tc.col])
			}
		})
	}
}

func Test3x3Matrix(t *testing.T) {
	mat := [][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}
	testCases := []struct {
		row int
		col int
		expected float64
	}{
		{0, 0, -3},
		{0, 1, 5},
		{0, 2, 0},
		{1, 0, 1},
		{1, 1, -2},
		{1, 2, -7},
		{2, 0, 0},
		{2, 1, 1},
		{2, 2, 1},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("row=%d,col=%d", tc.row, tc.col), func(t *testing.T) {
			if mat[tc.row][tc.col] != tc.expected {
				t.Errorf("Expected %f, got %f", tc.expected, mat[tc.row][tc.col])
			}
		})
	}
}

func TestEqualMatrices(t *testing.T) {
	A := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	B := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	if !MatrixEqual(A, B) {
		t.Errorf("Expected %v, got %v", true, MatrixEqual(A, B))
	}
}

func TestDifferentMatrices(t *testing.T) {
	A := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	B := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 9, 7, 6},
		{5, 4, 3, 2},
	}
	if MatrixEqual(A, B) {
		t.Errorf("Expected %v, got %v", false, MatrixEqual(A, B))
	}
}

func TestMatrixMultiply(t *testing.T) {
	A := [][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	}
	B := [][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	}
	C := MatrixMultiply(A, B)
	expected := [][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	}
	if !MatrixEqual(C, expected) {
		t.Errorf("Expected %v, got %v", expected, C)
	}
}
