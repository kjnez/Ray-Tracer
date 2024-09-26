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
