package bootcamp

func FibonacciRecursion(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// func main() {
// 	fmt.Println(FibonacciRecursion(-1)) // -1
// 	fmt.Println(FibonacciRecursion(0))  // 0
// 	fmt.Println(FibonacciRecursion(1))  // 1
// 	fmt.Println(FibonacciRecursion(2))  // 1
// 	fmt.Println(FibonacciRecursion(3))  // 2
// 	fmt.Println(FibonacciRecursion(4))  // 3
// 	fmt.Println(FibonacciRecursion(5))  // 5
// 	fmt.Println(FibonacciRecursion(6))  // 8
// }
