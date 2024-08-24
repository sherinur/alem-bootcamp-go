package bootcamp

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	return rune(s[len(s)-1])
}

// func main() {
// 	ap.PutRune(LastRune("salem"))
// 	ap.PutRune(LastRune("student"))
// 	ap.PutRune('\n')
// }
