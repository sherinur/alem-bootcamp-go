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
	var l int
	fmt.Scanf("%d", &l)

	arr := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	// print
	for i, num := range arr {
		check(num * 2)
		if i == len(arr)-1 {
			return
		}
		ap.PutRune(' ')
	}
}
