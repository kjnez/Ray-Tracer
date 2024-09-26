package features

import (
	// "math"
)

func MatrixEqual(A, B [][]float64) bool {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		return false
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if !Equal(A[i][j], B[i][j]) {
				return false
			}
		}
	}
	return true
}

func MatrixMultiply(A, B [][]float64) [][]float64 {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		panic("Dimension unmatched")
	}
	M := make([][]float64, len(A))
	for i := range M {
		M[i] = make([]float64, len(A[0]))
	}
	for row := 0; row < len(A); row++ {
		for col := 0; col < len(A[0]); col++ {
			for i := 0; i < len(A[0]); i++ {
				M[row][col] += A[row][i] * B[i][col]
			}
		}
	}
	return M
}
