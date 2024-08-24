package bootcamp

func SliceMakeN(n int) []int {
	if n < 0 {
		return nil
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	return arr
}

// func main() {
// 	fmt.Println(SliceMakeN(5))  // [0, 1, 2, 3, 4]
// 	fmt.Println(SliceMakeN(10)) // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
// 	fmt.Println(SliceMakeN(0))
// 	fmt.Println(SliceMakeN(1))
// 	fmt.Println(SliceMakeN(-10))
// }
