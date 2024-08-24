package bootcamp

func MagicGrowthN(n int) []string {
	result := []string{}
	runes := []rune("0123456789")

	if n < 1 || n > len(runes) {
		return result
	}

	if n == 1 {
		for _, c := range runes {
			result = append(result, string(c))
		}
		return result
	}

	for _, v := range MagicGrowthN(n - 1) {
		for _, c := range runes {
			if c <= rune(v[len(v)-1]) {
				continue
			}
			result = append(result[:], v+string(c))
		}
	}

	return result
}

// func main() {
// 	fmt.Println(MagicGrowthN(1))  // ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
// 	fmt.Println(MagicGrowthN(2))  // ["01", "02", ... "08", "09", "12", "13" ... "78", "79", "89"]
// 	fmt.Println(MagicGrowthN(3))  // ["012", "013", ... "089", "123", "123" ... "678", "679", "789"]
// 	fmt.Println(MagicGrowthN(9))  // ["012345678", "012345679", "012345689", ..., "123456789"]
// 	fmt.Println(MagicGrowthN(10)) // ["0123456789"]
// }
