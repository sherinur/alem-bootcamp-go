package bootcamp

func MagicGrowth2() []string {
	var result []string

	for i := 0; i <= 8; i++ {
		for j := i + 1; j <= 9; j++ {
			result = append(result, string('0'+i)+string('0'+j))
		}
	}
	return result
}

// func main() {
// 	fmt.Println(MagicGrowth2()) // ["01", "02", ... "08", "09", "12", "13" ... "78", "79", "89"]
// }
