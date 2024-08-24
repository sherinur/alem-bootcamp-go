package bootcamp

func MapSet(m map[string]int, key string, value int) {
	m[key] = value
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	fmt.Println(m) // map[one:1 two:2 three:3]
// 	MapSet(m, "two", -2)
// 	MapSet(m, "four", 4)
// 	fmt.Println(m) // map[one:1 two:-2 three:3 four:4]
// }
