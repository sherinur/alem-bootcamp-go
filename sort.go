package bootcamp

func Sort(arr []int, fn func(int, int) bool) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if !fn(arr[i], arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

// func main() {
// 	ascending := func(a, b int) bool {
// 		return a < b
// 	}

// 	descending := func(a, b int) bool {
// 		return a > b
// 	}

// 	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
// 	Sort(arr, ascending)
// 	fmt.Println(arr) // [1 1 2 3 4 5 5 6 9]

// 	arr = []int{3, 1, 4, 1, 5, 9, 2, 6, 5}
// 	Sort(arr, descending)
// 	fmt.Println(arr) // [9 6 5 5 4 3 2 1 1]
// }
