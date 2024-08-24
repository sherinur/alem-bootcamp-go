package bootcamp

func CountSubstr(s string, substr string) int {
	substrLen := len(substr)
	sLen := len(s)
	if sLen == 0 || substrLen == 0 || sLen < substrLen {
		return 0
	}

	count := 0
	start := 0

	for start <= sLen-substrLen {
		if s[start:start+substrLen] == substr {
			count++
			start += substrLen
		} else {
			start++
		}
	}

	return count
}

// func main() {
// 	fmt.Println(CountSubstr("qanagattandyrylmagandyqtarynyzdan", "an"))
// }
