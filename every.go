package bootcamp

func Every(arr []int, fn func(int) bool) bool {
	for _, num := range arr {
		if fn(num) == false {
			return false
		}
	}
	return true
}

// func main() {
// 	isEven := func(n int) bool {
// 		return n%2 == 0
// 	}

// 	result := Every([]int{2, 4, 6, 8}, isEven)
// 	fmt.Println(result) // true

// 	result = Every([]int{2, 4, 5, 8}, isEven)
// 	fmt.Println(result) // false
// }
