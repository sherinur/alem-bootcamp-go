package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	fmt.Println("Enter text: Characters:", Characters(text))
	fmt.Println("Sentences:", Sentences(text))
	fmt.Println("Words:", Words(text))
	fmt.Println("Letters:", Letters(text))
	fmt.Println("Vowels:", Vowels(text))
	fmt.Println("Consonants:", Consonants(text))
	fmt.Println("Digits:", Digits(text))
	fmt.Println("Spaces:", Spaces(text))
	fmt.Println("Special Characters:", SpecialCharacters(text))
}

func Characters(s string) int {
	return len(s)
}

func Sentences(s string) int {
	temp := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '!' || s[i] == '?' {
			temp++
		}
	}
	if s[len(s)-1] == '.' || s[len(s)-1] == '!' || s[len(s)-1] == '?' {
		temp--
	}
	return temp
}

func Words(s string) int {
	c := 0
	inWord := false
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for _, char := range s {
		if strings.ContainsRune(letters, char) {
			if !inWord {
				c++
				inWord = true
			}
		} else if !strings.ContainsRune(alphanumeric, char) {
			inWord = false
		}
	}
	return c
}

func Letters(s string) int {
	temp := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			temp++
		}
	}
	return temp
}

func Vowels(s string) int {
	temp := 0
	vowels := "aeiouAEIOU"
	for _, r := range s {
		for _, i := range vowels {
			if r == i {
				temp++
			}
		}
	}
	return temp
}

func Consonants(s string) int {
	temp := 0
	consonants := "BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"
	for _, r := range s {
		for _, i := range consonants {
			if r == i {
				temp++
			}
		}
	}
	return temp
}

func Digits(s string) int {
	temp := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			temp++
		}
	}
	return temp
}

func Spaces(s string) int {
	temp := 0
	for _, r := range s {
		if r == ' ' {
			temp++
		}
	}
	return temp
}

func SpecialCharacters(s string) int {
	temp := 0
	special := `!"#$%&'()*+,-./:;<=>?@[\]^_` + "`" + `{|}~`
	for _, r := range s {
		for _, i := range special {
			if r == i {
				temp++
			}
		}
	}
	return temp
}

// Hello world. How are you today? I hope fine!
