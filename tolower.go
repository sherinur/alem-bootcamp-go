package bootcamp

func ToLower(s string) string {
	var result string = ""
	if len(s) == 0 {
		return result
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			result += string(s[i] + ('a' - 'A'))
		} else {
			result += string(s[i])
		}
	}

	return result
}

// func main() {
// 	fmt.Println(ToLower("salem "))        // "SALEM "
// 	fmt.Println(ToLower("Salem Student")) // "SALEM STUDENT"
// 	fmt.Println(ToLower("S4LEm"))         // "S4LEM"

// 	// fmt.Println(string('s' - 32))
// }
