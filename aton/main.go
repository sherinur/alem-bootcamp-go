package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func check(a int) {
	if a < 0 {
		ap.PutRune('-')
		a = -a
	}

	if a == 0 {
		ap.PutRune('0')
		return
	}

	printDigits(a)
}

func printDigits(a int) {
	if a == 0 {
		return
	}

	printDigits(a / 10)
	ap.PutRune('0' + rune(a%10))
}

func main() {
	var n int
	var line string

	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s", &line)

	if len(line) == 0 {
		return
	}

	for i := 0; i < n && i < len(line); i++ {
		check(int(line[i]))
		if i != n-1 && i != len(line)-1 {
			ap.PutRune(' ')
		}
	}
}
