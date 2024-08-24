package bootcamp

func ArraySum(arr [20]int) int {
	sum := 0
	for _, element := range arr {
		sum += element
	}

	return sum
}

// func main() {
// 	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
// 	sum := ArraySum(arr)
// 	// 210
// }
