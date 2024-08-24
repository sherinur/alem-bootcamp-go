package bootcamp

func ArrayEqual(arr1, arr2 *[20]int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// func main() {
// 	arr1 := [20]int{77, 69, 76, 65}
// 	arr2 := [20]int{77, 69, 76, 65}
// 	fmt.Println(ArrayEqual(&arr1, &arr2))            // true
// 	fmt.Println(ArrayEqual(&arr1, &[20]int{77, 78})) // false
// }
