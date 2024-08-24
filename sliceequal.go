package bootcamp

func SliceEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) || cap(arr1) != cap(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// func main() {
// 	arr1 := []int{77, 69, 76, 65}
// 	arr2 := []int{77, 69, 76, 65}
// 	fmt.Println(SliceEqual(arr1, arr2))          // true
// 	fmt.Println(SliceEqual(arr1, []int{77, 78})) // false
// }
