package bootcamp

func ContainsSlice(slice []int, num int) bool {
	if len(slice) == 0 {
		return false
	}

	for _, value := range slice {
		if value == num {
			return true
		}
	}
	return false
}

func SliceUnion(slices ...[]int) []int {
	var result []int
	for _, slice := range slices {
		for _, num := range slice {
			if ContainsSlice(result, num) {
				continue
			} else {
				result = append(result, num)
			}
		}
	}
	return result
}

// func main() {
// 	result := SliceUnion([]int{1, 2, 3, 20}, []int{1, 2, 3, 4}, []int{15, 0, 2})
// 	fmt.Println(result) // [1, 2, 3, 20, 4, 5, 15, 0]
// }
