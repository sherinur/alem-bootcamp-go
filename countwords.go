package bootcamp

func splitToLowerWords(s string) []string {
	var words []string

	word := ""
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == ' ' {
			words = append(words, word)
			word = ""
			continue
		}
		if char >= 'A' && char <= 'Z' {
			char += 32
		}
		if char >= 'a' && char <= 'z' {
			word += string(char)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return words
}

func CountWords(s string) map[string]int {
	words := splitToLowerWords(s)
	numberOfWords := make(map[string]int)

	for _, word := range words {
		numberOfWords[word]++
	}

	return numberOfWords
}

// func main() {
// 	s := "The soup was stirred and stirred until thickened."
// 	wordCounts := CountWords(s)
// 	fmt.Println(wordCounts) // map[the:1 soup:1 was:1 and:1 stirred:2 until:1 thickened:1]
// }
