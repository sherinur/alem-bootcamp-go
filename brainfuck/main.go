package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func interpret(code string) {
	memory := make([]byte, 30000)
	ptr := 0
	codeIndex := 0
	loopStack := []int{}

	for codeIndex < len(code) {
		switch code[codeIndex] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			Print(string(memory[ptr]))
		case ',':
			if len(os.Args) > 1 {

				input := os.Args[1]

				if len(input) > 0 {
					memory[ptr] = input[0]
				} else {
					memory[ptr] = 0
				}

			} else {
				memory[ptr] = 0
			}
		case '[':
			if memory[ptr] == 0 {
				loopCount := 1
				for loopCount > 0 {
					codeIndex++
					if code[codeIndex] == '[' {
						loopCount++
					} else if code[codeIndex] == ']' {
						loopCount--
					}
				}
			} else {
				loopStack = append(loopStack, codeIndex)
			}
		case ']':
			if memory[ptr] != 0 {
				codeIndex = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}

		codeIndex++
	}
}

func Print(s string) {
	for _, v := range s {
		ap.PutRune(rune(v))
	}
}

func main() {
	args := os.Args
	interpret(args[1])
}
