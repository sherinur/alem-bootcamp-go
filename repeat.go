package bootcamp

func Repeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

// func main() {
// 	fmt.Println(Repeat("abc", 3)) // "abcabcabc"
// 	fmt.Println(Repeat("123", 2)) // "123123"
// }
