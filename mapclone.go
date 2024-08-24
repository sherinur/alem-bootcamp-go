package bootcamp

func MapClone(m map[string]int) map[string]int {
	clone := make(map[string]int)
	for key, value := range m {
		clone[key] = value
	}

	return clone
}

// func main() {
// 	m := map[string]int{"one": 1, "two": 2, "three": 3}
// 	clone := MapClone(m)
// 	fmt.Println(clone) // map[one:1 two:2 three:3]
// 	clone["four"] = 4
// 	fmt.Println(clone) // map[one:1 two:2 three:3 four:4]
// 	fmt.Println(m)     // map[one:1 two:2 three:3]
// }
