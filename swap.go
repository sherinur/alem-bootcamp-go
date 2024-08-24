package bootcamp

func Swap(a, b *int) {
	if a == nil || b == nil {
		return
	}
	*a, *b = *b, *a
}

// func main() {
// 	x := 1
// 	y := 2

// 	Swap(&x, &y)

// 	fmt.Println("x:", x, "y:", y)
// 	// Output:
// 	// x: 2 y: 1
// }
