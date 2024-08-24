package bootcamp

func RepeatEcho(s string) string {
	var result string
	var stack []string

	for _, c := range s {
		if c != ')' {
			stack = append(stack, string(c))
			continue
		} else {
			subStr := ""
			for stack[len(stack)-1] != "(" {
				subStr = stack[len(stack)-1] + subStr
				stack = stack[0 : len(stack)-1]
			}
			stack = stack[0 : len(stack)-1]

			k := ""
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				k = stack[len(stack)-1] + k
				stack = stack[0 : len(stack)-1]
			}

			num, _ := atoi(k)
			subStr = repeat(subStr, num)

			stack = append(stack, subStr)
		}
	}

	for _, s := range stack {
		result += s
	}

	return result
}

func atoi(input string) (int, string) {
	var n int
	for i, b := range []byte(input) {
		b -= '0'
		if b > 9 {
			return n, input[i:]
		}
		n = n*10 + int(b)
	}
	return n, ""
}

func repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
