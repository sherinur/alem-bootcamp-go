package main

import (
	"fmt"
	"math/rand"

	"github.com/alem-platform/ap"
)

func putString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func randomNum() int {
	return rand.Intn(100)
}

func main() {
	randomNumber := randomNum()

	for {
		var choice int
		putString("Guess number: ")
		fmt.Scanf("%d", &choice)
		if choice > randomNumber {
			putString("Lower\n")
		} else if choice < randomNumber {
			putString("Higher\n")
		} else {
			putString("Match, you win!\n")
			return
		}
	}
}
