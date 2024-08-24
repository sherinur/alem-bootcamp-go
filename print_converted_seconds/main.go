package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args[1]
	num := atoi(args)
	arr := make([]int, 3)
	day := 0
	if num > 2147483647 || num < 0 {
		ap.PutRune('N')
		ap.PutRune('V')
		ap.PutRune('\n')
	} else {
		for i := 0; i < 3; i++ {
			arr[i] = num % 60
			if i != 2 {
				num = num / 60
			}
		}
		day = num / 24
		arr[2] = arr[2] % 24
		if day != 0 {
			putnumber(day)
			ap.PutRune('d')
			ap.PutRune(' ')
		}
		if arr[2] != 0 {
			putnumber(arr[2])
			ap.PutRune('h')
			ap.PutRune(' ')
		}
		if arr[1] != 0 {
			putnumber(arr[1])
			ap.PutRune('m')
			ap.PutRune(' ')
		}

		if arr[0] != 0 {
			putnumber(arr[0])
			ap.PutRune('s')
		}
		ap.PutRune('\n')
	}
}

func atoi(s string) int {
	t := true
	n := 0
	if s[0] == '-' || s[0] == '+' || (s[0] >= 48 && s[0] <= 57) {
		for i := 1; i < len(s); i++ {
			if s[i] >= 48 && s[i] <= 57 {
			} else {
				t = false
			}
		}
	} else {
		t = false
	}

	if t == false {
		return n
	} else {
		if s[0] == '-' {
			m := 1
			for i := len(s) - 1; i > 0; i-- {
				n += m * (int(s[i]) - 48)
				m *= 10
			}
			n = 0 - n
		} else if s[0] == '+' {
			m := 1
			for i := len(s) - 1; i > 0; i-- {
				n += m * (int(s[i]) - 48)
				m *= 10
			}
		} else {
			m := 1
			for i := len(s) - 1; i >= 0; i-- {
				n += m * (int(s[i]) - 48)
				m *= 10
			}
		}
		return n
	}
}

func putnumber(n int) {
	digits := []rune("0123456789")
	if n < 0 {
		ap.PutRune('-')
		n *= -1
	}

	if n < 10 {
		ap.PutRune(digits[n])
		return
	}

	putnumber(n / 10)
	ap.PutRune(digits[n%10])
}
