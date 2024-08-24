package bootcamp

func Capitalize(s string) string {
	lowerS := ""
	var result string = ""

	// toLower
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			lowerS += string(s[i] + ('a' - 'A'))
		} else {
			lowerS += string(s[i])
		}
	}

	// first letter
	if len(lowerS) > 0 && lowerS[0] >= 'a' && lowerS[0] <= 'z' {
		result += string(lowerS[0] - ('a' - 'A'))
	} else if len(lowerS) > 0 {
		result += string(lowerS[0])
	}

	// other words
	for i := 1; i < len(lowerS); i++ {
		if lowerS[i-1] == ' ' && lowerS[i] >= 'a' && lowerS[i] <= 'z' {
			result += string(lowerS[i] - ('a' - 'A'))
		} else {
			result += string(lowerS[i])
		}
	}

	return result
}

// func main() {
// 	fmt.Println(Capitalize("salem student"))
// 	fmt.Println(Capitalize("SALEM-5TUDENT, Q41aisyn?"))
// 	fmt.Println(Capitalize("&mIJ"))
// }
