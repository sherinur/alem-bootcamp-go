package bootcamp

func MissingNumber(arr []int) int {
	positiveSet := make(map[int]bool)

	for _, num := range arr {
		if num > 0 {
			positiveSet[num] = true
		}
	}

	i := 1
	for {
		if !positiveSet[i] {
			return i
		}
		i++
	}
}

// func main() {
// 	fmt.Println(MissingNumber([]int{1, 2, 3, 4, 5}))   // 6
// 	fmt.Println(MissingNumber([]int{1, 2, 4, 5}))      // 3
// 	fmt.Println(MissingNumber([]int{3, 4, 5, 6, 7}))   // 1
// 	fmt.Println(MissingNumber([]int{}))                // 1
// 	fmt.Println(MissingNumber([]int{0, -1, -2, 1, 2})) // 3
// }
