package bootcamp

func SlicePushFront(arr *[]int, values ...int) {
	lengthOfPushing := len(values)
	*arr = append(*arr, values...)

	for i := len(*arr) - 1 - lengthOfPushing; i >= 0; i-- {
		(*arr)[i+lengthOfPushing] = (*arr)[i]
	}

	for i := range values {
		(*arr)[i] = values[i]
	}
}

// func main() {
// 	arr := []int{1}
// 	fmt.Println(arr) // [1]

// 	SlicePushFront(&arr, 10)
// 	fmt.Println(arr) // [10, 1]

// 	SlicePushFront(&arr, 20)
// 	fmt.Println(arr) // [20, 10, 1]

// 	SlicePushFront(&arr, 5, 3)
// 	fmt.Println(arr) // [5, 3, 20, 10, 1]

// 	// arr := []int{1, 2, 3, 4, 5}
// 	// // 9 8 7
// 	// fmt.Println(arr)
// 	// arr = append(arr, 9, 8, 7)

// 	// for i := len(arr) - 4; i >= 0; i-- {
// 	// 	arr[i+3] = arr[i]
// 	// }

// 	// fmt.Println(arr)
// }
