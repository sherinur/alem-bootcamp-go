package main

import (
	"github.com/alem-platform/ap"
)

func main() {
	lower := 97
	upper := 65

	for i := 0; i < 26; i++ {
		ap.PutRune(rune(lower))
		ap.PutRune(rune(upper))
		lower++
		upper++
	}

	ap.PutRune('\n')
}
