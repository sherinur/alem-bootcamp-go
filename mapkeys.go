package bootcamp

func MapKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	m2 := map[string]int{"RbgMY": 210, "gQhQCp": 193, "hIUr": 94, "homGab": 234, "qqA": 135, "x": 0}
// 	keys := MapKeys(m)
// 	fmt.Println(keys)        // [one two three]
// 	fmt.Println(MapKeys(m2)) // [one two three]
// }
