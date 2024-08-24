package bootcamp

// resheto Eratosphes
func getPrimes(n int) []int {
	if n < 2 {
		return []int{}
	}
	// generating an array
	primes := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		primes[i] = i + 2
	}
	primesLen := len(primes)
	// erathosfenes
	for i := 0; i < primesLen; i++ {
		if primes[i] != 0 {
			for j := i + 1; j < len(primes); j++ {
				if primes[j]%primes[i] == 0 {
					primes[j] = 0
				}
			}
		}
	}

	result := []int{}
	for _, value := range primes {
		if value != 0 {
			result = append(result, value)
		}
	}

	return result
}

func NextPrime(n int) int {
	if n < 0 || n == 0 {
		return 2
	}

	primes := getPrimes(n * 2)
	prime := 0

	for _, value := range primes {
		if value > n {
			prime = value
			break
		}
	}
	return prime
}

// func main() {
// 	fmt.Println(NextPrime(10)) // 11
// 	fmt.Println(NextPrime(11)) // 13
// 	fmt.Println(NextPrime(-1)) // 2
// }
