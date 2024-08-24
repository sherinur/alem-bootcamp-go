package bootcamp

import (
	"github.com/alem-platform/ap"
)

func PutNumber(n int) {
	if n == 0 {
		ap.PutRune('0')
		return
	}

	if n < 0 {
		ap.PutRune('-')
		n = -n
	}

	var digits []int
	for n != 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] >= 0 && digits[i] <= 9 {
			ap.PutRune('0' + rune(digits[i]))
		}
	}
}
