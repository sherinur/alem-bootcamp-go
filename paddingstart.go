package bootcamp

func PaddingStart(s string, totalLength int) string {
	lens := RuneCountInString(s)
	if totalLength <= lens {
		return s
	}
	str := ""
	for i := 0; i < totalLength-lens; i++ {
		str += " "
	}
	str += s
	return str
}

func RuneCountInString(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
