package bootcamp

func Replace(s string, old string, new string) string {
	if len(s) == 0 {
		return ""
	}

	lenOld := len(old)

	var result string
	start := 0
	for start <= len(s)-lenOld {
		if s[start:start+lenOld] == old {
			result += new
			start += lenOld
		} else {
			result += string(s[start])
			start++
		}
	}

	result += s[start:]
	return result
}

// func main() {
// 	fmt.Println(Replace("Hello student!", "Hello", "Salem"))
// 	fmt.Println(Replace("banana", "a", "o"))
// }
