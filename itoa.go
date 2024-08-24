package bootcamp

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	signFlag := false
	if n < 0 {
		signFlag = true
		n *= -1
	}

	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}

	if signFlag {
		result = "-" + result
	}

	return result
}

// func main() {
// 	fmt.Println(Itoa(123))
// 	fmt.Println(Itoa(-456))
// 	fmt.Println(Itoa(0))
// }
