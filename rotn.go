package bootcamp

func RotN(s string, n int) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	n = n % 26

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			if (char + rune(n)) > 'z' {
				result += string('a' + (char + rune(n) - 'z' - 1))
			} else {
				result += string(char + rune(n))
			}
		} else if char >= 'A' && char <= 'Z' {
			if (char + rune(n)) > 'Z' {
				result += string('A' + (char + rune(n) - 'Z' - 1))
			} else {
				result += string(char + rune(n))
			}
		} else {
			result += string(char)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(RotN("SALEM", 1))
// 	fmt.Println(RotN("abc", 13))
// }
