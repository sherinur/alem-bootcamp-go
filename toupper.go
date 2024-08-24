package bootcamp

func ToUpper(s string) string {
	var result string = ""
	if len(s) == 0 {
		return result
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			result += string(s[i] - ('a' - 'A'))
		} else {
			result += string(s[i])
		}
	}

	return result
}

// func main() {
// 	fmt.Println(ToUpper("salem "))        // "SALEM "
// 	fmt.Println(ToUpper("Salem Student")) // "SALEM STUDENT"
// 	fmt.Println(ToUpper("S4LEm"))         // "S4LEM"

// 	// fmt.Println(string('s' - 32))
// }
