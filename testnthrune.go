package bootcamp

func TestNthRune(fn func(s string, n int) rune) bool {
	tests := []struct {
		s        string
		n        int
		expected rune
	}{
		{"hello", 1, 'h'},
		{"hello", 5, 'o'},
		{"hello", 0, -1},
		{"hello", 6, -1},
		{"", 1, -1},
		{"你好", 1, '你'},
		{"こんにちは", 3, 'に'},
	}

	for _, test := range tests {
		result := fn(test.s, test.n)
		if result != test.expected {
			return false
		}
	}

	return true
}
