package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	processName, err := os.Executable()
	if err != nil {
		return
	}
	result := ""
	for i := len(processName) - 1; i >= 0; i-- {
		if processName[i] == '/' {
			break
		}
		result += string(processName[i])
	}

	for i := len(result) - 1; i >= 0; i-- {
		ap.PutRune(rune(result[i]))
	}

	ap.PutRune('\n')
}
