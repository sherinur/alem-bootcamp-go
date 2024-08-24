package bootcamp

func ArraySlice(arr [20]int, low int, high int) []int {
	// check for negative values and invalid values
	if low < high && low >= 0 && high <= len(arr) {
		return arr[low:high]
	}

	return nil
}

// func main() {
// 	arr := [20]int{
// 		67, 40, 114, 142, 258, 156, 288, 71, 250, 183, 24, 173, 222, 58, 34, 268, 134, 35, 48, 121,
// 	}
// 	fmt.Println(ArraySlice(arr, 10, 20))
// 	// [24, 173, 222, 58, 34, 268, 134, 35, 48, 121]
// 	// fmt.Println(ArraySlice(arr, 5, 5))  // []
// 	// fmt.Println(ArraySlice(arr, 5, 1))  // []
// 	// fmt.Println(ArraySlice(arr, -1, 1)) // []
// 	// fmt.Println(ArraySlice(arr, 0, 0))  // []
// 	// fmt.Println(ArraySlice(arr, 0, 21)) // []
// }
