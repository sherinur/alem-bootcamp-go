package bootcamp

func MergeSortedArrays(arr1, arr2 []int) []int {
	i := 0
	j := 0
	result := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

// func main() {
// 	arr1 := []int{1, 3, 5, 7}
// 	arr2 := []int{2, 4, 6, 8}
// 	result := MergeSortedArrays(arr1, arr2)
// 	fmt.Println(result) // [1 2 3 4 5 6 7 8]

// 	arr1 = []int{0, 2, 4}
// 	arr2 = []int{1, 3, 5}
// 	result = MergeSortedArrays(arr1, arr2)
// 	fmt.Println(result) // [0 1 2 3 4 5]

// 	arr1 = []int{10, 20, 30}
// 	arr2 = []int{5, 15, 25, 35}
// 	result = MergeSortedArrays(arr1, arr2)
// 	fmt.Println(result) // [5 10 15 20 25 30 35]
// }
