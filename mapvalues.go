package bootcamp

func MapValues(m map[string]int) []int {
	values := make([]int, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	values := MapValues(m)
// 	fmt.Println(values) // [1 2 3]
// }
