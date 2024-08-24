package bootcamp

func ConvertNbrBase(n int, base string) string {
	var arr string
	divider := len(base)
	temp := true
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				temp = false
			}
		}
	}
	if divider < 2 {
	} else if n < 0 {
	} else if temp == false {
	} else {
		for n > 0 {
			remains := n % divider
			arr = string(base[remains]) + arr
			n = n / divider
		}
	}

	return arr
}

// func main() {
// 	result := ConvertNbrBase(-1465, "0123456789")
// 	fmt.Println(result) // 1465

// 	result = ConvertNbrBase(-1465, "01")
// 	fmt.Println(result) // 10110111001

// 	result = ConvertNbrBase(1465, "01234567")
// 	fmt.Println(result) // 2671

// 	result = ConvertNbrBase(0, "qwertyu")
// 	fmt.Println(result) // 5B9

// 	result = ConvertNbrBase(614, "lEI7eAC4xvgWDjyXGuUo8isF52tZh")
// 	fmt.Println(result) //
// }
