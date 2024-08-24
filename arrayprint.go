package bootcamp

import "github.com/alem-platform/ap"

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

func ArrayPrint(arr [20]int) {
	for i, element := range arr {
		check(element)

		if i == len(arr)-1 {
			return
		}

		ap.PutRune(' ')
	}
}

// func main() {
// 	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
// 	ArrayPrint(arr)
// 	// Output:
// 	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
// }
