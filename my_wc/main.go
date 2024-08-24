package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 {
		printString("usage: my_wc [-lwc] file\n")
		return
	}

	var countLines, countWords, countBytes bool
	var fileName string

	for _, arg := range arg {
		switch arg {
		case "-l":
			countLines = true
		case "-w":
			countWords = true
		case "-c":
			countBytes = true
		default:
			fileName = arg
		}
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		printString("my_wc: " + fileName + ": No such file or directory\n")
		return
	}

	lines, words, bytes := countFile(file)

	if !countLines && !countWords && !countBytes {
		countLines = true
		countWords = true
		countBytes = true
	}

	if countLines {
		putNumber(lines)
		printString(" ")
	}
	if countWords {
		putNumber(words)
		printString(" ")
	}
	if countBytes {
		putNumber(bytes)
		printString(" ")
	}

	printString(fileName + "\n")
}

func countFile(file []byte) (int, int, int) {
	lines := 0
	words := 0

	for _, ch := range file {
		if ch == '\n' {
			lines++
		}
	}
	if len(file) > 0 && file[len(file)-1] != '\n' {
		lines++
	}
	wordsF := false
	for _, char := range file {
		if char == ' ' {
			if wordsF {
				words++
				wordsF = false
			}
		} else {
			wordsF = true
		}
	}

	if wordsF {
		words++
	}
	return lines, words, len(file)
}

func printString(s string) {
	for _, r := range s {
		ap.PutRune(r)
	}
}

func putNumber(n int) {
	switch {
	case n > 0:
		a := n % 10
		n /= 10
		if n == 0 {
			ap.PutRune(rune(a + 48))
			return
		}
		putNumber(n)
		ap.PutRune(rune(a + 48))
	case n < 0:
		ap.PutRune('-')
		putNumber(-n)
	default:
		ap.PutRune('0')
	}
}
