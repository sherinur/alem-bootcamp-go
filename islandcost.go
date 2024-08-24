package bootcamp

func CostPoint(matrix *[][]int, cost *int, x, y int) {
	if y >= 0 && y < len((*matrix)) && x >= 0 && x < len((*matrix)[y]) {
		if (*matrix)[y][x] != 0 {
			*cost += (*matrix)[y][x]
			(*matrix)[y][x] = 0
			CostPoint(matrix, cost, x, y+1)
			CostPoint(matrix, cost, x, y-1)
			CostPoint(matrix, cost, x+1, y)
			CostPoint(matrix, cost, x-1, y)
		}
	}
}

func IslandCost(matrix [][]int, x, y int) int {
	cost := 0
	CostPoint(&matrix, &cost, x, y)
	return cost
}
