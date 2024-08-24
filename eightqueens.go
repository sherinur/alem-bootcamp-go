package bootcamp

type Point struct {
	x int
	y int
}

var results = make([][]int, 0)

func EightQueens() [][]int {
	for col := 0; col < 8; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, 8)
	}
	return results
}

func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			var result []int
			for _, point := range current {
				result = append(result, point.x+1)
			}
			results = append(results, result)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y + 1; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, append([]Point{}, current...), n)
				}
			}
		}
	}
}

func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	dx := a.x - b.x
	if dx < 0 {
		dx = -dx
	}
	dy := a.y - b.y
	if dy < 0 {
		dy = -dy
	}
	return a.x == b.x || a.y == b.y || dx == dy
}
