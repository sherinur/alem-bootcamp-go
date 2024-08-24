package bootcamp

func Eraser(s string) string {
	stack := []rune{}

	for _, char := range s {
		if char == '<' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, char)
		}
	}

	return string(stack)
}

// func main() {
// 	fmt.Printf("%q\n", Eraser("s<alem"))          // "alem"
// 	fmt.Printf("%q\n", Eraser("s<alem<<<"))       // "a"
// 	fmt.Printf("%q\n", Eraser("s<alem<< <<? D<")) // "a? "
// 	fmt.Printf("%q\n", Eraser("<<<"))             // ""
// }
