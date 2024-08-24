package bootcamp

import "github.com/alem-platform/ap"

func Sign(n int) {
	if n > 0 {
		ap.PutRune('+')
	} else if n < 0 {
		ap.PutRune('-')
	} else {
		ap.PutRune('0')
	}
}

// func main() {
//     Sign(2024)
// }
