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
