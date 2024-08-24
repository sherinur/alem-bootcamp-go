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
	var first, second int

	fmt.Scanf("%d %d", &first, &second)

	sum := first + second
	sub := first - second
	multi := first * second

	check(sum)
	ap.PutRune(' ')
	check(sub)
	ap.PutRune(' ')
	check(multi)
	ap.PutRune(' ')

	if second != 0 {
		div := first / second
		check(div)
	} else {
		ap.PutRune('F')
	}
}
