// import "fmt"
package bootcamp

func SliceBatch(slice []int, size int) [][]int {
	if size <= 0 {
		return nil
	}

	var mainSlice [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size

		if end > len(slice) {
			end = len(slice)
		}
		mainSlice = append(mainSlice, slice[i:end])
	}
	return mainSlice
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7}
// 	batch := SliceBatch(arr, 2)
// 	for _, v := range batch {
// 		fmt.Println(v)
// 	}
// 	// Output:
// 	// [1, 2]
// 	// [3, 4]
// 	// [5, 6]
// 	// [7]