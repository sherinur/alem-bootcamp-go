package bootcamp

func IncrementIt(n *int) {
	if n == nil {
		return
	} else {
		*n += 1
	}
}

// func main() {
// 	num := 10
// 	IncrementIt(&num)
// 	fmt.Println(num) // 11
// 	IncrementIt(&num)
// 	fmt.Println(num) // 12
// }
