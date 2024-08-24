package bootcamp

import "github.com/alem-platform/ap"

func PrintBooks(books []*Book) {
	maxName := 0
	maxAuthor := 0

	for i := 0; i < len(books); i++ {
		if len(books[i].Name) > maxName {
			maxName = len(books[i].Name)
		}

		if len(books[i].Author) > maxAuthor {
			maxAuthor = len(books[i].Author)
		}
	}

	putString("Name ")
	putSpaces(maxName - 4)

	putString("Author ")
	putSpaces(maxAuthor - 6)

	putString("Year")

	ap.PutRune('\n')

	for i := 0; i < len(books); i++ {
		putString(books[i].Name + " ")
		putSpaces(maxName - len(books[i].Name))

		putString(books[i].Author + " ")
		putSpaces(maxAuthor - len(books[i].Author))

		putString(intToString(books[i].Year))

		ap.PutRune('\n')
	}
}

func putSpaces(number int) {
	for i := 0; i < number; i++ {
		ap.PutRune(' ')
	}
}

func intToString(num int) string {
	if num == 0 {
		return "0"
	}

	var result string
	negative := false

	if num < 0 {
		negative = true
		num *= -1
	}

	for num > 0 {
		digit := num % 10
		result = string('0'+digit) + result
		num /= 10
	}

	if negative {
		result = "-" + result
	}

	return result
}

func putString(line string) {
	for _, c := range line {
		ap.PutRune(c)
	}
}

// func main() {
// 	books := []*Book{
// 		{Name: "The Kaizen Way", Author: "Robert Maurer", Year: 2009},
// 		{Name: "Dialogs", Author: "Plato", Year: -400},
// 		{Name: "Unknown", Author: "Unknown", Year: 10},
// 	}

// 	PrintBooks(books)
// }
