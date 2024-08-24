package bootcamp

func Split(s string, sep string) []string {
	if sep == "" {
		var chars []string
		for _, char := range s {
			chars = append(chars, string(char))
		}
		return chars
	}

	var result []string
	start := 0
	sepLen := len(sep)

	for start <= len(s) {
		index := start
		found := false

		for index+sepLen <= len(s) && !found {
			if s[index:index+sepLen] == sep {
				found = true
			} else {
				index++
			}
		}

		if found {
			result = append(result, s[start:index])
			start = index + sepLen
		} else {
			result = append(result, s[start:])
			break
		}
	}

	if len(s) >= sepLen && s[len(s)-sepLen:] == sep {
		result = append(result, "")
	}

	return result
}

// func main() {
// 	fmt.Println(Split("a,b,c", ","))
// 	fmt.Println(Split("salem-student", "-"))
// 	fmt.Println(Split("salem", ""))
// 	fmt.Println(Split("tw$t%whLT?8*IL0ThTT*hT^c", "T"))
// }
