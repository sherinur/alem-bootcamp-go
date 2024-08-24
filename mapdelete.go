package bootcamp

func MapDelete(m map[string]int, key string) {
	delete(m, key)
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	fmt.Println(m) // map[one:1 two:2 three:3]
// 	MapDelete(m, "two")
// 	fmt.Println(m) // map[one:1 three:3]
// 	MapDelete(m, "four")
// 	fmt.Println(m) // map[one:1 three:3]
// }
