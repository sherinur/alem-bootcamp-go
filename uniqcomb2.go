package bootcamp

func UniqComb2(characters string) []string {
	result := []string{}
	runes := []rune(characters)

	if len(runes) < 2 {
		return []string{}
	}

	seen := make(map[rune]bool)
	for _, c := range runes {
		if seen[c] {
			return []string{}
		}
		seen[c] = true
	}

	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes); j++ {
			if i != j {
				result = append(result, string(runes[i])+string(runes[j]))
			}
		}
	}

	return result
}

// func main() {
// 	fmt.Println(UniqComb2("abc")) // ["ab", "ac", "ba", "bc", "ca", "cb"]
// 	fmt.Println(UniqComb2("ab"))  // ["ab", "ba"]
// 	fmt.Println(UniqComb2("a"))   // []
// 	fmt.Println(UniqComb2("aba")) // []
// }
