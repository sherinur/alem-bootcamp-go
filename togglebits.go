package bootcamp

func ToggleBits(n byte) byte {
	return n ^ 255
}

// func main() {
// 	var b byte = 255
// 	fmt.Printf("%08b\n", b) // 11111111
// 	b = ToggleBits(b)
// 	fmt.Printf("%08b\n", b) // 00000000
// }
