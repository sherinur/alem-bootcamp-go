package bootcamp

func SliceDelete(arr *[]int, low, high int) bool {
	if low < 0 || high < 0 || low >= high || high > len(*arr) || low > len(*arr) {
		return false
	}
	*arr = append((*arr)[:low], (*arr)[high:]...)

	return true
}

// func main() {
// 	arr := []int{0, 1, 2, 3, 4, 5}

// 	fmt.Println(arr)                     // [0, 1, 2, 3, 4, 5]
// 	fmt.Println(SliceDelete(&arr, 1, 3)) // true
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 3, 3)) // false
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 5, 3)) // false
// 	fmt.Println(arr)                     // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, -10, 5)) // false
// 	fmt.Println(arr)                       // [0, 3, 4, 5]

// 	fmt.Println(SliceDelete(&arr, 0, 4)) // true
// 	fmt.Println(arr)                     // []
// }
