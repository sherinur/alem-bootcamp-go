package bootcamp

// import "fmt"

func CompareThem(a, b *int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		return false
	} else {
		return *a == *b
	}
}

// func main() {
// 	var first int = 10
// 	var second int = 11
// 	var third int = 10

// 	fmt.Println(CompareThem(&first, &second)) // false
// 	fmt.Println(CompareThem(&first, &third))  // true
// }
