package bootcamp

func PowIterative(x int, power int) int {
	if power < 0 {
		return -1
	}

	result := 1
	for i := 1; i <= power; i++ {
		result *= x
	}

	return result
}

// func main() {
// 	fmt.Println(PowIterative(2, -1)) // -1
// 	fmt.Println(PowIterative(2, 0))  // 1
// 	fmt.Println(PowIterative(2, 1))  // 2
// 	fmt.Println(PowIterative(2, 2))  // 4
// 	fmt.Println(PowIterative(2, 3))  // 8
// 	fmt.Println(PowIterative(2, 4))  // 16
// }
