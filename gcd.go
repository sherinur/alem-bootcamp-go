package bootcamp

func GCD(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	rem := a % b
	rem = GCD(b, a%b)
	return rem
}
