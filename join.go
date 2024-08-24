package bootcamp

func Join(elements []string, sep string) string {
	if len(elements) <= 0 {
		return ""
	}
	result := ""

	for i, element := range elements {
		result += element
		if i != len(elements)-1 {
			result += sep
		}
	}
	return result
}

// func main() {
// 	fmt.Println(Join([]string{"salem", "student"}, " "))
// 	fmt.Println(Join([]string{"1", "2", "3"}, ", "))
// }
