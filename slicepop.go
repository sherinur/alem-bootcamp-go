package bootcamp

func SlicePop(arr *[]int) int {
	if len(*arr) == 0 {
		return 0
	}

	deleted := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]

	return deleted
}

// func main() {
// 	arr := []int{5, 10, 20}
// 	size := len(arr)

// 	for i := 0; i < size; i++ {
// 		deleted := SlicePop(&arr)
// 		fmt.Println(deleted)
// 	}
// 	// Output:
// 	// 20
// 	// 10
// 	// 5

// 	deleted := SlicePop(&arr) // 0
// 	fmt.Println(deleted)
// }
