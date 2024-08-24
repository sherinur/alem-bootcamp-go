package bootcamp

func Rot13(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""

	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			if char <= 'm' {
				result += string(char + 13)
			} else {
				result += string(char - 13)
			}
		} else if char >= 'A' && char <= 'Z' {
			if char <= 'M' {
				result += string(char + 13)
			} else {
				result += string(char - 13)
			}
		} else {
			result += string(char)
		}
	}

	return result
}

// func main() {
// 	fmt.Println(Rot13("salem"))
// 	fmt.Println(Rot13("fnyrz"))
// }
