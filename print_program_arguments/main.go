package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		for _, char := range arg {
			ap.PutRune(char)
		}
		ap.PutRune('\n')
	}
}
