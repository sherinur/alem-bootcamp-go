package bootcamp

func RomanToInt(s string) int {
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		char := string(s[i])
		_, ok := roman[char]
		if !ok {
			return 0
		}

		value := roman[char]
		if value >= prevValue {
			result += value
		} else {
			result -= value
		}
		prevValue = value
	}

	return result
}

// func main() {
// 	fmt.Println(RomanToInt("III"))     // 3
// 	fmt.Println(RomanToInt("IV"))      // 4
// 	fmt.Println(RomanToInt("IX"))      // 9
// 	fmt.Println(RomanToInt("LVIII"))   // 58
// 	fmt.Println(RomanToInt("MCMXCIV")) // 1994
// 	fmt.Println(RomanToInt(""))        // 0
// 	fmt.Println(RomanToInt("salem"))   // 0
// 	fmt.Println(RomanToInt("IAMGAY"))  // 0
// }
