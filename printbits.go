package bootcamp

import (
	"github.com/alem-platform/ap"
)

func printL(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func PrintBits(n byte) {
	binary := ""
	for i := 7; i >= 0; i-- {
		bit := (n >> uint(i)) & 1
		binary += string('0' + byte(bit))
	}

	if len(binary) < 8 {
		for i := 0; i < 8-len(binary); i++ {
			binary = "0" + binary
		}
	}

	printL(binary)
	ap.PutRune('\n')
}

// func main() {
// 	PrintBits(5)
// 	// 00000101
// 	PrintBits(255)
// 	// 11111111
// 	PrintBits(0)
// 	// 00000000
// 	PrintBits(101)
// }
