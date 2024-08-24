package bootcamp

func IsNumeric(s string) bool {
	signCount := 0
	for i := 0; i < len(s); i++ {
		if signCount > 1 {
			return false
		}

		if s[i] == 45 || s[i] == 43 {
			signCount++
			continue
		}
		isNumber := s[i] >= 48 && s[i] <= 57

		if !isNumber {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(IsNumericSimple("123"))     // true
// 	fmt.Println(IsNumericSimple("-123"))    // true
// 	fmt.Println(IsNumericSimple("+123"))    // true
// 	fmt.Println(IsNumericSimple("+-123"))   // false
// 	fmt.Println(IsNumericSimple(" 123"))    // false
// 	fmt.Println(IsNumericSimple("123 abc")) // false
// 	fmt.Println(IsNumericSimple("123abc"))  // false
// 	fmt.Println(IsNumericSimple("ab"))      // false
// }
