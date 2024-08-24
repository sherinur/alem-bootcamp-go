package bootcamp

func ArraySetOne(arr *[20]int, idx int, val int) bool {
	if idx > cap(arr) || idx < 0 {
		return false
	} else {
		arr[idx] = val
	}

	return true
}

// func main() {
// 	arr := [20]int{1, 2, 3}
// 	ok := ArraySetOne(&arr, -9, 5)
// 	fmt.Println("ok:", ok)   // ok: true
// 	fmt.Println("arr:", arr) // arr: [1, 5, 3, ...]
// }
