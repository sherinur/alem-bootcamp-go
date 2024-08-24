package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func intToString(i int) string {
	if i == 0 {
		return "0"
	}
	var neg bool
	if i < 0 {
		neg = true
		i = -i
	}
	var res []byte
	for i > 0 {
		res = append(res, byte(i%10)+'0')
		i /= 10
	}
	if neg {
		res = append(res, '-')
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

func main() {
	arguments := len(os.Args) - 1

	s := intToString(arguments)
	for _, c := range s {
		ap.PutRune(rune(c))
	}

	ap.PutRune('\n')
}
