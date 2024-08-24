package bootcamp

func SliceClone(src, dst *[]int) {
	clone := make([]int, len(*src), cap(*src))
	copy(clone, *src)
	*dst = clone
}

// func main() {
// 	arr := make([]int, 10)
// 	for i, v := range []int{10, 20, 13, 5, 12, 31} {
// 		arr[i] = v
// 	}

// 	clone := []int{}

// 	fmt.Println("arr:", arr, len(arr), cap(arr))         // arr: [10, 20, 13, 5, 12, 31] 6 10
// 	fmt.Println("clone:", clone, len(clone), cap(clone)) // clone: [] 0 0

// 	SliceClone(&arr, &clone)

// 	fmt.Println("arr:", arr, len(arr), cap(arr))         // arr: [10, 20, 13, 5, 12, 31] 6 10
// 	fmt.Println("clone:", clone, len(clone), cap(clone)) // clone: [10, 20, 13, 5, 12, 31] 6 10
// }
