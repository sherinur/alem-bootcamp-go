package bootcamp

func Glob(s string, pattern string) bool {
	sIndex, patternIndex := 0, 0
	sLen, patternLen := len(s), len(pattern)
	sStar, pStar := -1, -1

	for sIndex < sLen {
		if patternIndex < patternLen && pattern[patternIndex] == '?' {
			sIndex++
			patternIndex++
		} else if patternIndex < patternLen && pattern[patternIndex] == '*' {
			pStar = patternIndex
			sStar = sIndex
			patternIndex++
		} else if patternIndex < patternLen && pattern[patternIndex] == '[' {
			patternIndex++
			negate := false
			if patternIndex < patternLen && (pattern[patternIndex] == '!' || pattern[patternIndex] == '^') {
				negate = true
				patternIndex++
			}
			match := false
			for patternIndex < patternLen && pattern[patternIndex] != ']' {
				if patternIndex+1 < patternLen && pattern[patternIndex+1] == '-' {
					start := pattern[patternIndex]
					end := pattern[patternIndex+2]
					if (s[sIndex] >= start && s[sIndex] <= end) || (s[sIndex] >= end && s[sIndex] <= start) {
						match = true
					}
					patternIndex += 3
				} else {
					if s[sIndex] == pattern[patternIndex] {
						match = true
					}
					patternIndex++
				}
			}
			if (match && !negate) || (!match && negate) {
				sIndex++
				patternIndex++
			} else {
				return false
			}
		} else if patternIndex < patternLen && s[sIndex] == pattern[patternIndex] {
			sIndex++
			patternIndex++
		} else if pStar != -1 {
			patternIndex = pStar + 1
			sStar++
			sIndex = sStar
		} else {
			return false
		}
	}

	for patternIndex < patternLen && pattern[patternIndex] == '*' {
		patternIndex++
	}

	return patternIndex == patternLen
}
