package bootcamp

func ReverseBits(n byte) byte {
	var num byte = 0
	for i := 0; i < 8; i++ {
		num += (n >> i) & 1 << (7 - i)
	}
	return num
}
