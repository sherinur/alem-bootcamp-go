package bootcamp

func IsNil(a *int) bool {
	if a == nil {
		return true
	} else {
		return false
	}
}

// func main() {
// 	var a int = 10
// 	var b *int
// 	fmt.Println(a, IsNil(&a))
// 	fmt.Println(b, IsNil(b))
// 	fmt.Println(nil, IsNil(nil))
// }
