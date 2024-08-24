package bootcamp

func DigitsOnly(s string) string {
	var digits string
	bytes := []byte(s)

	for _, b := range bytes {
		if b >= '0' && b <= '9' {
			digits += string(b)
		}
	}
	return digits
}

// func main() {
// 	fmt.Println(DigitsOnly("abc123"))
// 	fmt.Println(DigitsOnly("Sa1em student! 23"))
// 	fmt.Println(DigitsOnly("no digits here!"))
// }
