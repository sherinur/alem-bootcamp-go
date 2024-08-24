package bootcamp

func FibonacciIterative(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	}
	fib1 := 1
	fib2 := 1
	i := 0
	for i < n-2 {
		fib_sum := fib1 + fib2
		fib1 = fib2
		fib2 = fib_sum
		i += 1
	}

	return fib2
}

// func main() {
// 	fmt.Println(FibonacciIterative(-1)) // -1
// 	fmt.Println(FibonacciIterative(0))  // 0
// 	fmt.Println(FibonacciIterative(1))  // 1
// 	fmt.Println(FibonacciIterative(2))  // 1
// 	fmt.Println(FibonacciIterative(3))  // 2
// 	fmt.Println(FibonacciIterative(4))  // 3
// 	fmt.Println(FibonacciIterative(5))  // 5
// 	fmt.Println(FibonacciIterative(6))  // 8
// }
