package bootcamp

func copyMatrix(dst, src [][]rune) {
	for i := range dst {
		copy(dst[i], src[i])
	}
}

func SliceMatrixRotateN(matrix [][]rune, turns int) {
	n := len(matrix)
	if turns%4 == 0 || n == 0 || len(matrix[0]) != n {
		return
	}

	turns = (turns%4 + 4) % 4

	for t := 0; t < turns; t++ {
		rotatedMatrix := make([][]rune, n)
		for i := range rotatedMatrix {
			rotatedMatrix[i] = make([]rune, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				rotatedMatrix[j][n-1-i] = matrix[i][j]
			}
		}
		copyMatrix(matrix, rotatedMatrix)
	}
}

// func main() {
// 	var matrix = [][]rune{
// 	  { 'a', 'b', 'c' },
// 	  { 'd', 'e', 'f' },
// 	  { 'g', 'h', 'i' },
// 	}

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// a b c
// 	// d e f
// 	// g h i

// 	SliceMatrixRotateN(matrix, 1)

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// g d a
// 	// h e b
// 	// i f c

// 	SliceMatrixRotateN(matrix, -2)

// 	SliceMatrixPrint(matrix)
// 	// Output:
// 	// c f i
// 	// b e h
// 	// a d g
//   }
