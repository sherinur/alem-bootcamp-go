package bootcamp

func Some(arr []int, fn func(int) bool) bool {
	for _, num := range arr {
		if fn(num) {
			return true
		}
	}

	return false
}

// func main() {
// 	isEven := func(n int) bool {
// 		return n%2 == 0
// 	}

// 	result := Some([]int{1, 3, 5, 7, 8}, isEven)
// 	fmt.Println(result) // true

// 	result = Some([]int{1, 3, 5, 7}, isEven)
// 	fmt.Println(result) // false
// }
