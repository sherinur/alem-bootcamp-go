package bootcamp

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	MergeSort(arr[:len(arr)/2])
	MergeSort(arr[len(arr)/2:])
	merged := merge(arr[:len(arr)/2], arr[len(arr)/2:])
	copy(arr, merged)
}

// func main() {
// 	arr := []int{38, 27, 43, 3, 9, 82, 10}
// 	MergeSort(arr)
// 	fmt.Println(arr)
// 	// Output: [3 9 10 27 38 43 82]
// }
