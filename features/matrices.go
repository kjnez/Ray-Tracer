package features

// "math"

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

func MatrixTupleMultiply(A [][]float64, b Tuple) Tuple {
	if len(A[0]) != 4 {
		panic("Columns of A don't match size of b")
	}
	c := Tuple{}
	c.x = A[0][0] * b.x + A[0][1] * b.y + A[0][2] * b.z + A[0][3] * b.w
	c.y = A[1][0] * b.x + A[1][1] * b.y + A[1][2] * b.z + A[1][3] * b.w
	c.z = A[2][0] * b.x + A[2][1] * b.y + A[2][2] * b.z + A[2][3] * b.w
	c.w = A[3][0] * b.x + A[3][1] * b.y + A[3][2] * b.z + A[3][3] * b.w
	return c
}

func MatrixTranspose(A [][]float64) [][]float64 {
	result := make([][]float64, len(A[0]))
	for i := range result {
		result[i] = make([]float64, len(A))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			result[j][i] = A[i][j]
		}
	}
	return result
}
