package bootcamp

func Trim(s string) string {
	if len(s) == 0 {
		return s
	}

	var result string

	// leading spaces
	leadingCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			leadingCount++
		} else {
			break
		}
	}
	result = s[leadingCount:len(s)]

	// trailing spaces
	trailingCount := 0
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ' ' {
			trailingCount++
		} else {
			break
		}
	}
	result = result[:len(result)-trailingCount]
	return result
}

// func main() {
// 	fmt.Println(Trim("   Salem student!   "))
// 	fmt.Println(Trim("   Trim spaces   "))
// }
