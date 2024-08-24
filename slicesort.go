package bootcamp

func SliceSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // swap
			}
		}
	}
}

// func main() {
// 	arr := []int{10, 90, 20, 5, 12, 3, 0}
// 	SliceSort(arr)
// 	fmt.Println(arr) // [0, 3, 5, 10, 12, 20, 90]
// }
