package bootcamp

func IsPrime(n int) bool {
	if n < 0 || n == 0 || n == 1 {
		return false
	} else {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				return false
			}
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsPrime(997)) // true
// 	fmt.Println(IsPrime(4))   // false
// 	fmt.Println(IsPrime(-11)) // false
// }
