package bootcamp

func SliceConcat(slices ...[]int) []int {
	arr := []int{}

	for _, i := range slices {
		for _, j := range i {
			arr = append(arr, j)
		}
	}

	return arr
}

// func main() {
// 	result := SliceConcat([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, []int{15, 0, 2})
// 	fmt.Println(result) // [1, 2, 3, 4, 5, 1, 2, 3, 4, 15, 0, 2]
// }
