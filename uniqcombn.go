package bootcamp

func UniqCombN(characters string, n int) []string {
	result := []string{}
	runes := []rune(characters)

	if len(runes) < n {
		return []string{}
	}

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes); j++ {
			if i != j {
				if runes[i] == runes[j] {
					return []string{}
				}
			}
		}
	}

	if n == 1 {
		for _, c := range runes {
			result = append(result, string(c))
		}
		return result
	}

	for i := 0; i < len(runes); i++ {
		char := string(runes[i])

		remainingChars := characters[:i] + characters[i+1:]
		remainingCombinations := UniqCombN(remainingChars, n-1)

		for _, comb := range remainingCombinations {
			result = append(result, char+comb)
		}
	}

	return result
}

// func main() {
// 	fmt.Println(UniqCombN("abc", 1))  // ["a", "b", "c"]
// 	fmt.Println(UniqCombN("abc", 2))  // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqCombN("ab", 2))   // ["ab", "ba"]
// 	fmt.Println(UniqCombN("a", 1))    // ["a"]
// 	fmt.Println(UniqCombN("ab", 3))   // []
// 	fmt.Println(UniqCombN("aa", 1))   // []
// }
