package bootcamp

func SliceAppendN(n int) []int {
	if n < 0 {
		return nil
	}

	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}

	return arr
}

// func main() {
// 	fmt.Println(SliceAppendN(5))  // [0, 1, 2, 3, 4]
// 	fmt.Println(SliceAppendN(10)) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// 	fmt.Println(SliceAppendN(0))
// 	fmt.Println(SliceAppendN(-1))
// 	fmt.Println(SliceAppendN(1))
// }
