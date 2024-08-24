package bootcamp

func ConvertBaseNbr(n string, base string) int {
	result := 0
	mult := len(base)
	temp := true
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				temp = false
				break
			}
		}
	}
	reverse := ""
	for k := len(n) - 1; k >= 0; k-- {
		reverse += string(n[k])
	}
	k := 0
	for i := 0; i < len(reverse); i++ {
		for j := 0; j < len(base); j++ {
			if reverse[i] == base[j] {
				k++
			}
		}
	}
	if k != len(reverse) {
		temp = false
	}
	if temp {
		t := 0
		for i := 0; i < len(reverse); i++ {
			for j := 0; j < len(base); j++ {
				if reverse[i] == base[j] {
					t = j
				}
			}
			result += t * pow(mult, i)
			t = 0
		}
	} else {
		return -1
	}
	return result
}

// func main() {
//  fmt.Println(ConvertBaseNbr("1465B", "0123456789"))     // 1465
//  fmt.Println(ConvertBaseNbr("10110111009", "01"))       // 1465
//  fmt.Println(ConvertBaseNbr("2671", "01234567"))        // 1465
//  fmt.Println(ConvertBaseNbr("5B9", "0123456789ABCDEF")) // 1465
//  fmt.Println(ConvertBaseNbr("1", "00"))                 // -1
// }

func pow(x int, power int) int {
	if power < 0 {
		return -1
	} else if power == 0 {
		return 1
	} else {
		return x * pow(x, power-1)
	}
}
