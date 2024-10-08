package bootcamp

func SliceGetClone(src []int) []int {
	clone := make([]int, len(src), cap(src))
	copy(clone, src)

	return clone
}

// func main() {
// 	arr := make([]int, 10)
// 	for i, v := range []int{10, 20, 13, 5, 12, 31} {
// 		arr[i] = v
// 	}

// 	fmt.Println("arr:", arr, len(arr), cap(arr)) // arr: [10, 20, 13, 5, 12, 31] 6 10

// 	clone := SliceGetClone(arr)

// 	fmt.Println("arr:", arr, len(arr), cap(arr))         // arr: [10, 20, 13, 5, 12, 31] 6 10
// 	fmt.Println("clone:", clone, len(clone), cap(clone)) // clone: [10, 20, 13, 5, 12, 31] 6 10

// 	clone[0] = 0

// 	fmt.Println("arr:", arr, len(arr), cap(arr))         // arr: [0, 20, 13, 5, 12, 31] 6 10
// 	fmt.Println("clone:", clone, len(clone), cap(clone)) // clone: [10, 20, 13, 5, 12, 31] 6 10
// }
