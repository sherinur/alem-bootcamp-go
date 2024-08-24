package bootcamp

func RotateWords(s string) string {
	n := len(s)
	var words []string
	var spaces []string
	var currentWord []byte
	var currentSpaces []byte

	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []byte{}
			}
			currentSpaces = append(currentSpaces, s[i])
		} else {
			if len(currentSpaces) > 0 {
				spaces = append(spaces, string(currentSpaces))
				currentSpaces = []byte{}
			}
			currentWord = append(currentWord, s[i])
		}
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}
	if len(currentSpaces) > 0 {
		spaces = append(spaces, string(currentSpaces))
	}

	if len(words) > 1 {
		lastWord := words[len(words)-1]
		copy(words[1:], words[:len(words)-1])
		words[0] = lastWord
	}

	result := ""
	wordIndex := 0
	spaceIndex := 0

	for wordIndex < len(words) || spaceIndex < len(spaces) {
		if wordIndex < len(words) {
			result += words[wordIndex]
			wordIndex++
		}
		if spaceIndex < len(spaces) {
			result += spaces[spaceIndex]
			spaceIndex++
		}
	}

	return result
}
