package bootcamp

func PowRecursion(x int, power int) int {
	if power < 0 {
		return -1
	} else if x == 0 {
		return 1
	}

	result := 1
	if power <= 0 {
		return result
	}
	return x * PowRecursion(x, power-1)
}

// func main() {
// 	fmt.Println(PowRecursion(2, -1)) // -1
// 	fmt.Println(PowRecursion(2, 0))  // 1
// 	fmt.Println(PowRecursion(2, 1))  // 2
// 	fmt.Println(PowRecursion(2, 2))  // 4
// 	fmt.Println(PowRecursion(2, 3))  // 8
// 	fmt.Println(PowRecursion(2, 4))  // 16
// 	fmt.Println(PowRecursion(0, 5))
// 	fmt.Println(PowRecursion(2, 10))
// }
