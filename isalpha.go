package bootcamp

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		isInInterval := (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122)
		if !isInInterval {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("HelloWorld"))  // true
// 	fmt.Println(IsAlpha("Hello123"))    // false
// 	fmt.Println(IsAlpha("Hello World")) // false

// 	// for i := range "1234567890" {
// 	// 	fmt.Println(i)
// 	// }

// 	// fmt.Printf("%T", "1")
// 	// fmt.Println("wad9"[3])
// 	// fmt.Printf("%T", "1234")
// 	// fmt.Printf("%T", "gEJ4"[3])
// }
