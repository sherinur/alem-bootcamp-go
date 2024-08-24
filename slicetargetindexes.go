package bootcamp

func SliceTargetIndexes(arr []int, target int) []int {
	targetIndexes := []int{}

	for i := range arr {
		if arr[i] == target {
			targetIndexes = append(targetIndexes, i)
		}
	}

	return targetIndexes
}

// func main() {
// 	arr := []int{2, 3, 2, 5, 6, 7, 8, 2, 10}
// 	fmt.Println(SliceTargetIndexes(arr, 2)) // [0, 2, 7]
// 	fmt.Println(SliceTargetIndexes(arr, 1)) // []
// }
