package bootcamp

func IsPalindromeWord(s string) bool {
	sLower := ""

	if len(s) == 0 {
		return false
	} else {
		for i := 0; i < len(s); i++ {
			if s[i] >= 'A' && s[i] <= 'Z' {
				sLower += string(s[i] + ('a' - 'A'))
			} else {
				sLower += string(s[i])
			}
		}
	}

	left := 0
	right := len(s) - 1
	i := 0
	for left <= right {
		isInInterval := (sLower[i] >= 65 && sLower[i] <= 90) || (sLower[i] >= 97 && sLower[i] <= 122)
		if !isInInterval {
			return false
		}

		if sLower[left] == sLower[right] {
			left++
			right--
			i++
		} else {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPalindromeWord("racecar")) // true
// 	fmt.Println(IsPalindromeWord("level"))   // true
// 	fmt.Println(IsPalindromeWord("salem"))   // false
// 	fmt.Println(IsPalindromeWord("saas"))    // true
// 	fmt.Println(IsPalindromeWord("_asa_"))   // false
// 	fmt.Println(IsPalindromeWord("12a21"))   // false
// 	fmt.Println(IsPalindromeWord("QazAq"))   // true
// }
