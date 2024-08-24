package bootcamp

func MapContains(m map[string]int, key string) bool {
	_, ok := m[key]
	return ok
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	fmt.Println(MapContains(m, "two"))  // true
// 	fmt.Println(MapContains(m, "four")) // false
// }
