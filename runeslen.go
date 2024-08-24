package bootcamp

func RunesLen(arr [20]rune) int {
	for i, element := range arr {
		if element == '\x00' {
			return i
		}
	}

	return len(arr)
}

// func main() {
// 	arr := [20]rune{'a', 'b', 'c', '0', 'e'}
// 	length := RunesLen(arr)
// 	fmt.Println(length)
// 	// 3
// }
