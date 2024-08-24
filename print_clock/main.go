package main

import "github.com/alem-platform/ap"

func check(a int) {
	if a < 10 {
		ap.PutRune('0')
		ap.PutRune('0' + rune(a))
	} else {
		ap.PutRune('0' + rune(a/10))
		ap.PutRune('0' + rune(a%10))
	}
}

func main() {
	for hour := 0; hour < 24; hour++ {
		for minute := 0; minute < 60; minute++ {
			check(hour)
			ap.PutRune(':')
			check(minute)
			ap.PutRune('\n')
		}
	}
}
