package main

import (
	"os"

	"github.com/alem-platform/ap"
)

var dogMaps = map[string]*[]string{
	"simple":   &simple,
	"maltipoo": &maltipoo,
	"tazy":     &tazy,
	"cur":      &cur,
}

var simple = []string{
	"  \\",
	"^..^      /",
	"/_/\\_____/",
	"   /\\   /\\",
	"  /  \\ /  \\",
}

var maltipoo = []string{
	"  \\",
	"   \\ __    __",
	"   o-''))_____\\\\",
	"   \"--__/ * * * )",
	"   c_c__/-c____/",
}

var tazy = []string{
	"  \\                __",
	"   \\___________   /  \\",
	"               \\ / ..|\\",
	"                (_\\  |_)",
	"                /  \\@'",
	"               /     \\",
	"           _  /  `   |",
	"          \\\\ /  \\  | _\\",
	"           \\   /_ || \\\\_",
	"            \\____)|_) \\_)",
}

var cur = []string{
	"   \\",
	"    \\ D\\___/\\",
	"     \\ (0_o)",
	"        (V)",
	"----oOo--U--oOo------",
	"_______|_______|_____",
}

func main() {
	args := os.Args[1:]
	key := "simple"
	var text string

	if len(args) > 3 || len(args) < 1 {
		printString("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
	} else if args[0] == "-b" {
		if len(args) > 2 {
			key = args[1]
			text = args[2]
		} else {
			printString("usage: dogsay [-b cur|maltipoo|simple|tazy] text\n")
			return
		}
	} else {
		text = args[0]
	}

	printBubble(text, key)
}

func printBubble(text string, key string) {
	txt := sliceText(text)
	rowLen := maxLineLen(txt)

	if rowLen < 4 {
		rowLen = 4
	}
	for i := 0; i < len(txt)+2; i++ {
		if i == 0 {
			printString(" ")
			printN('_', rowLen+2)
			printString("\n")
		} else if i == len(txt)+1 {
			if len(txt) == 0 {
				printString("< ")
				printString(" >\n")
			}
			printString(" ")
			printN('-', rowLen+2)
			printString("\n")
		} else {
			if len(txt) == 1 {
				printString("< ")
				printString(txt[i-1])
				printString(" >\n")
			} else {
				if i == 1 {
					printString("/ ")
					printString(txt[i-1])
					printN(' ', rowLen-len(txt[i-1])+1)
					printString("\\\n")
				} else if i == len(txt) {
					printString("\\ ")
					printString(txt[i-1])
					printN(' ', rowLen-len(txt[i-1])+1)
					printString("/\n")
				} else {
					printString("| ")
					printString(txt[i-1])
					printN(' ', rowLen-len(txt[i-1])+1)
					printString("|\n")
				}
			}
		}
	}

	Dog(key)
}

func printString(s string) {
	for i, c := range s {
		ap.PutRune(c)
		if c == '\n' && i != len(s)-1 {
			ap.PutRune(' ')
		}
	}
}

func printN(c rune, times int) {
	for i := 0; i < times; i++ {
		ap.PutRune(c)
	}
}

func maxLineLen(text []string) int {
	l := 0

	for _, ch := range text {
		if len(ch) > l {
			l = len(ch)
		}
	}

	return l
}

func sliceText(text string) []string {
	slice := []string{}

	line := ""
	for i, ch := range text {
		if ch == '\n' {
			slice = append(slice, line)
			line = ""
		} else {
			line += string(ch)
		}

		if i == len(text)-1 {
			slice = append(slice, line)
			line = ""
		}
	}

	return slice
}

func Dog(key string) {
	for _, l := range *dogMaps[key] {
		printString(l + "\n")
	}
}
