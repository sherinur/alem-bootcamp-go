package bootcamp

func MapGet(m map[string]int, key string) int {
	return m[key]
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	fmt.Println(MapGet(m, "two"))  // 2
// 	fmt.Println(MapGet(m, "four")) // 0
// }
