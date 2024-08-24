package bootcamp

func IslandRemove(matrix [][]int, x, y int) {
	if x < 0 || y < 0 || y >= len(matrix) {
		return
	}
	if x >= len(matrix[y]) {
		return
	}

	if matrix[y][x] < 1 {
		return
	}
	matrix[y][x] = 0
	var v int

	v = x - 1
	if v > -1 {
		IslandRemove(matrix, v, y)
	}
	v += 2
	if v < len(matrix[y]) {
		IslandRemove(matrix, v, y)
	}

	v = y - 1
	if v > -1 {
		IslandRemove(matrix, x, v)
	}
	v += 2
	if v < len(matrix) {
		IslandRemove(matrix, x, v)
	}
	return
}
