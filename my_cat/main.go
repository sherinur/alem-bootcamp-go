package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func putString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		putString("\nusage: my_cat file ...\n")
		return
	}

	for i := 1; i < len(args); i++ {
		fileData, err := os.ReadFile(args[i])
		if err != nil {
			putString("my_cat: " + args[i] + ": No such file or directory\n")
			return
		}
		putString(string(fileData))
	}
}
