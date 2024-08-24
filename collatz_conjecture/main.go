package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	input := os.Args[1]
	for _, el := range input {
		if el < '0' || el > '9' {
			return
		}
	}

	num := Atoi(input)
	Print(input)

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
		Print(itoa(num))
	}
}

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	isNegative := s[0] == '-'
	if isNegative {
		s = s[1:]
	}
	result := 0
	for _, r := range s {
		digit := int(r - '0')
		if digit < 0 || digit > 9 {
			return 0
		}
		result = result*10 + digit
	}
	if isNegative {
		result = -result
	}
	return result
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}
	var result string
	for num > 0 {
		digit := num % 10
		result = string('0'+digit) + result
		num /= 10
	}
	return sign + result
}

func Print(s string) {
	for _, el := range s {
		ap.PutRune(el)
	}
	ap.PutRune('\n')
}
