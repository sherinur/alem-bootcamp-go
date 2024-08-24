package bootcamp

func Atoi(s string) int {
	bitDepth := 1
	result := 0
	sign := 1

	// sign test
	if s[0] == '-' {
		sign *= -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	// zero start test
	for i := 0; i < len(s); i++ {
		if s[0] == '0' {
			s = s[1:]
		} else {
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if byte(s[i]) < 48 || byte(s[i]) > 57 {
			return 0
		}
		result += bitDepth * int((byte(s[i]) - 48))
		bitDepth *= 10
	}

	result *= sign
	return result
}

// func main() {
// 	fmt.Println(Atoi("123"))
// 	fmt.Println(Atoi("+123"))
// 	fmt.Println(Atoi("-123"))
// 	fmt.Println(Atoi("-123!"))
// 	fmt.Println(Atoi("abc"))
// 	fmt.Println(Atoi("000124"))
// }
