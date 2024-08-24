package bootcamp

func ForEach(arr []int, fn func(*int)) {
	for i := 0; i < len(arr); i++ {
		fn(&arr[i])
	}
}

// func main() {
// 	increment := func(n *int) {
// 		*n++
// 	}

// 	arr := []int{1, 2, 3, 4}
// 	ForEach(arr, increment)

// 	fmt.Println(arr) // [2 3 4 5]
// }
