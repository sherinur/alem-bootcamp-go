package bootcamp

func GetIncrementor(start int, step int) func() int {
	incrementor := func() int {
		start += step
		return start
	}

	return incrementor
}

// func main() {
// 	incrementor := GetIncrementor(0, 1)
// 	fmt.Println(incrementor()) // 1
// 	fmt.Println(incrementor()) // 2
// 	fmt.Println(incrementor()) // 3

// 	anotherIncrementor := GetIncrementor(10, 3)
// 	fmt.Println(anotherIncrementor()) // 13
// 	fmt.Println(anotherIncrementor()) // 16
// }
