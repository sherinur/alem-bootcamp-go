package bootcamp

import "github.com/alem-platform/ap"

func SliceMatrixPrint(matrix [][]rune) {
	lastRow := len(matrix) - 1
	for i, line := range matrix {
		lastChar := len(line) - 1
		for j, char := range line {
			ap.PutRune(char)
			if j != lastChar {
				ap.PutRune(' ')
			}
		}
		if i != lastRow {
			ap.PutRune('\n')
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
// }
