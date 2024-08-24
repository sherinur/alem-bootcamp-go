package main

import (
	"github.com/alem-platform/ap"
)

func main() {
	for i := 122; i >= 97; i-- {
		ap.PutRune(rune(i))
	}

	ap.PutRune('\n')
}
