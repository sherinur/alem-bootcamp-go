package bootcamp

func Concat(s ...string) string {
	result := ""
	for _, line := range s {
		result += line
	}

	return result
}

// func main() {
// 	fmt.Println(Concat("Salem", " ", "Student!"))
// 	fmt.Println(Concat("1", "2", "3", "4", "5"))
// }
