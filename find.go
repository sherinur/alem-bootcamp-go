package bootcamp

func Find(arr []int, fn func(int) bool) *int {
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) {
			return &arr[i]
		}
	}

	return nil
}

// func main() {
// 	isEven := func(n int) bool {
// 		return n%2 == 0
// 	}

// 	result := Find([]int{1, 3, 5, 8, 10}, isEven)
// 	if result != nil {
// 		fmt.Println(*result) // 8
// 	} else {
// 		fmt.Println("No match found")
// 	}

// 	result = Find([]int{1, 3, 5, 7}, isEven)
// 	if result != nil {
// 		fmt.Println(*result)
// 	} else {
// 		fmt.Println("No match found") // No match found
// 	}
// }
