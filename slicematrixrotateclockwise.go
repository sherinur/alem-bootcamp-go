package bootcamp

func SliceMatrixRotateClockwise(matrix [][]rune) {
	n := len(matrix)
	if n == 0 || len(matrix[0]) != n {
		return
	}

	rotatedMatrix := make([][]rune, n)
	for i := range rotatedMatrix {
		rotatedMatrix[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotatedMatrix[j][n-1-i] = matrix[i][j]
		}
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = rotatedMatrix[i][j]
		}
	}
}

// func main() {
// 	matrix := [][]rune{
// 		{'a', 'b', 'c'},
// 		{'d', 'e', 'f'},
// 		{'g', 'h', 'i'},
// 	}

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// a b c
// 	// d e f
// 	// g h i

// 	SliceMatrixRotateClockwise(matrix)

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c
// }
