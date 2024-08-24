package bootcamp

func IsDone(matrix [][]int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func solve_KnightTour(matrix *[][][]int, temp [][]int, x int, y int, step int) {
	temp[x][y] = step

	moves := [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}

	for _, move := range moves {
		newX, newY := x+move[0], y+move[1]
		if newX >= 0 && newX < len(temp) && newY >= 0 && newY < len(temp[0]) && temp[newX][newY] == 0 {
			solve_KnightTour(matrix, copyMatrixInt(temp), newX, newY, step+1)
		}
	}

	if IsDone(temp) {
		*matrix = append(*matrix, temp)
		return
	}
}

func copyMatrixInt(matrix [][]int) [][]int {
	copyMat := make([][]int, len(matrix))
	for i := range matrix {
		copyMat[i] = make([]int, len(matrix[i]))
		copy(copyMat[i], matrix[i])
	}
	return copyMat
}

func KnightTour(size int) [][][]int {
	if size < 1 {
		return [][][]int{}
	}
	answer := make([][][]int, 0)
	temp := make([][]int, size)
	for i := range temp {
		temp[i] = make([]int, size)
	}
	solve_KnightTour(&answer, temp, 0, 0, 1)
	return answer
}
