package bootcamp

func Contains(str string, substr string) bool {
	if len(substr) == 0 || ((len(str) == len(substr)) && (str == substr)) {
		return true
	} else if len(substr) > len(str) {
		return false
	}

	start := 0
	for start <= len(str)-len(substr) {
		if str[start:start+len(substr)] == substr {
			return true
		} else {
			start++
		}
	}
	return false
}

// func main() {
// 	fmt.Println(Contains("hello world", "world"))
// 	fmt.Println(Contains("test", "best"))
// }
