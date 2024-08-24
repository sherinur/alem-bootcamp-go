package bootcamp

func _IslandExtend(matrix *[][]int, x, y int) bool {
	if y >= 0 && y < len((*matrix)) && x >= 0 && x < len((*matrix)[y]) {
		if (*matrix)[y][x] > 0 {
			(*matrix)[y][x] = -1
			_IslandExtend(matrix, x, y+1)
			_IslandExtend(matrix, x, y-1)
			_IslandExtend(matrix, x+1, y)
			_IslandExtend(matrix, x-1, y)
			return true
		} else {
			(*matrix)[y][x] = -1
		}
	}
	return false
}

func IslandExtend(matrix [][]int, x, y int) bool {
	newMap := make([][]int, len(matrix))
	for i := range newMap {
		newMap[i] = make([]int, len(matrix[0]))
	}
	for i, w := range matrix {
		for j, _ := range w {
			newMap[i][j] = matrix[i][j]
		}
	}
	val := _IslandExtend(&newMap, x, y)
	if val {
		val = false
		for i, w := range matrix {
			for j, _ := range w {
				if matrix[i][j] == 0 && newMap[i][j] == -1 {
					matrix[i][j] = 1
					val = true
				}
			}
		}
	}
	return val
}
