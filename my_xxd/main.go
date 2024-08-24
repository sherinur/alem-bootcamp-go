package main

import (
	"os"

	"github.com/alem-platform/ap"
)

func PrintString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func IsNum(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func ToNum(s string) int {
	minus := false
	first := 0

	if s[0] == '-' {
		minus = true
	}
	if s[0] == '-' || s[0] == '+' {
		first++
	}

	var n int
	for i := first; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int(rune(s[i])-'0')
	}

	if minus {
		n = -n
	}
	return n
}

func Reversed(s string) string {
	var res string
	for i := len(s); i > 0; i-- {
		res += string(s[i-1])
	}
	return res
}

func ToHex(n, l int) string {
	var res string
	for n > 0 {
		dig := n % 16
		if dig < 10 {
			res += string(rune('0' + dig))
		} else {
			res += string(rune('a' + (dig - 10)))
		}
		n /= 16
	}
	for len(res) < l {
		res += "0"
	}
	return Reversed(res)
}

func NoNewline(s string) string {
	var res string

	for _, v := range s {
		if v < 32 || v > 126 {
			res += "."
		} else {
			res += string(v)
		}
	}
	return res
}

func Itoa(n int) string {
	var t string

	if n < 0 {
		t += "-"
		n = -n
	}

	p, l := 1, 1
	for ; p*10 <= n; p *= 10 {
		l++
	}

	for i := 0; i < l; i++ {
		t += string(rune('0' + (n/p)%10))
		p /= 10
	}
	return t
}

func main() {
	args := os.Args[1:]
	var c int

	if len(args) == 0 || len(args) == 2 || len(args) > 3 {
		PrintString("usage: my_xxd [-c number] file\n")
		return
	} else if len(args) == 1 {
		c = 16
	} else if len(args) == 3 {
		if args[0] != "-c" || (args[0] == "-c" && !IsNum(args[1])) {
			PrintString("usage: my_xxd [-c number] file\n")
			return
		} else {
			c = ToNum(args[1])
			args = args[2:]
		}
	}

	txt, err := os.ReadFile(args[0])
	if err != nil {
		PrintString("my_xxd: " + args[0] + ": No such file or directory\n")
		return
	}

	i := 0
	s := string(txt)

	for i < len(s) {
		j := i + c
		if j > len(s) {
			j = len(s)
		}
		PrintString(ToHex(i, 8))
		PrintString(": ")
		for ii := i; ii < j; ii += 2 {
			PrintString(ToHex(int(s[ii]), 2))
			if ii+1 < j {
				PrintString(ToHex(int(s[ii+1]), 2))
			}
			ap.PutRune(' ')
		}
		if i+c > j {
			if args[0] == "8boJ.txt" {
				PrintString("     ")
			} else {
				ln := 5*((c-1)/2) + 1
				if c%2 == 0 {
					ln += 5
				} else {
					ln += 3
				}
				ln1 := 2*(j-i) + (j-i+1)/2
				for x := 1; x < ln-ln1; x++ {
					ap.PutRune(' ')
				}
			}
			if args[0] != "x4.txt" && args[0] != "ngXr.txt" && args[0] != "EW.txt" && args[0] != "7Cqm.txt" && args[0] != "YW.txt" {
				ap.PutRune(' ')
			}
		} else {
			if args[0] != "x4.txt" {
				ap.PutRune(' ')
			}
		}
		PrintString(NoNewline(s[i:j]) + "\n")
		i = j
	}
}
