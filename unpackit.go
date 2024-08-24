package bootcamp

// import "fmt"

func UnpackIt(p *******int) int {
	if p == nil || *p == nil || **p == nil || ***p == nil || ****p == nil || *****p == nil || ******p == nil {
		return 0
	} else {
		return *******p
	}
}

// func main() {
// 	var num int = 42
// 	var p1 *int = &num
// 	var p2 **int = &p1
// 	var p3 ***int = &p2
// 	var p4 ****int = &p3
// 	var p5 *****int = &p4
// 	var p6 ******int = &p5
// 	var p7 *******int = &p6

// 	fmt.Println(UnpackIt(p7)) // 42

// 	var nilPointer *******int
// 	fmt.Println(UnpackIt(nilPointer)) // 0
// }
