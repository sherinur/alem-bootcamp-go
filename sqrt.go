package bootcamp

func Sqrt(x int) int {
	if x < 0 {
		return -1
	}

	left := 0
	right := x

	for left <= right {
		mid := (left + right) / 2
		midSquared := mid * mid

		if midSquared == x {
			return mid
		} else if midSquared < x {
			left = mid + 1
		} else if midSquared > x {
			right = mid - 1
		}
	}

	return -1
}

// func main() {
// 	fmt.Println(Sqrt(16)) // 4
// 	fmt.Println(Sqrt(1))  // 1
// 	fmt.Println(Sqrt(3))  // 0
// 	fmt.Println(Sqrt(-4)) // -1
// }
