package bootcamp

// import "fmt"

func RoundBrackets(s string) bool {
	if len(s) == 0 {
		return false
	}

	var temp []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			temp = append(temp, s[i])
		} else if s[i] == ')' {
			if len(temp) == 0 {
				return false
			}
			temp = temp[:len(temp)-1]
		}
	}

	return len(temp) == 0
}

// func main() {
// 	fmt.Println(RoundBrackets("()()()"))   // true
// 	fmt.Println(RoundBrackets("(salem)"))  // true
// 	fmt.Println(RoundBrackets("(salem("))  // false
// 	fmt.Println(RoundBrackets(")(1)(f)(")) // false
// 	fmt.Println(RoundBrackets("))(()"))    // false
// 	fmt.Println(RoundBrackets(""))         // false
// }
