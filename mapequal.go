package bootcamp

func MapEqual(m1 map[string]int, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	} else {
		for key := range m1 {
			if m1[key] != m2[key] {
				return false
			}
		}
	}

	return true
}

// func main() {
// 	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
// 	m2 := map[string]int{"one": 1, "two": 2, "three": 3}
// 	m3 := map[string]int{"one": 1, "two": 2, "four": 4}

// 	fmt.Println(MapEqual(m1, m2)) // true
// 	fmt.Println(MapEqual(m1, m3)) // false
// }
