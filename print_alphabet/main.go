package main

import (
	"github.com/alem-platform/ap"
)

func main() {
	for i := 97; i < 123; i++ {
		ap.PutRune(rune(i))
	}

	ap.PutRune('\n')
}
