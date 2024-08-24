package bootcamp

import (
	"bufio"

	// "fmt"
	"os"
)

type Book struct {
	Name   string
	Author string
	Year   int
}

func separate(s string) []string {
	var words []string
	word := ""
	for j := 0; j < len(s); j++ {
		if s[j] != ',' {
			word += string(s[j])
		} else {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)
	}
	return words
}

func toInt(s string) int {
	num := 0
	sign := 1
	for i := 0; i < len(s); i++ {
		if i == 0 && s[i] == '-' {
			sign = -1
			continue
		}

		if s[i] > '9' || s[i] < '0' {
			break
		}

		digit := int(s[i] - '0')
		num = num*10 + digit
	}
	return num * sign
}

func GetBooksFromCsv(path string) []*Book {
	// open the file
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	// new scanner bufio
	scanner := bufio.NewScanner(file)
	var lines []string

	// storing scanned lines into lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// cut the header of the csv
	head := separate(lines[0])
	lines = lines[1:len(lines)]

	// by the head of the csv, defining the order of columns and storing in map
	order := make(map[string]int)

	for i, column := range head {
		order[column] = i
	}
	// deleting the garbage
	head = nil
	scanner = nil

	// new slice books
	books := []*Book{}

	for i := 0; i < len(lines); i++ {
		separatedLine := separate(lines[i])
		book := &Book{}

		book.Name = separatedLine[order["Name"]]
		book.Author = separatedLine[order["Author"]]
		book.Year = toInt(separatedLine[order["Year"]])
		books = append(books, book)
	}

	return books
}

// func main() {
// 	books := GetBooksFromCsv("books.csv")
// 	for _, b := range books {
// 		fmt.Println(b.Name, b.Author, b.Year)
// 	}

// 	books2 := GetBooksFromCsv("books2.csv")
// 	for _, b := range books2 {
// 		fmt.Println(b.Name, b.Author, b.Year)
// 	}

// 	GetBooksFromCsv("book.csv")
// }
