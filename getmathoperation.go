package bootcamp

func GetMathOperation(op string) *func(int, int) int {
	var operation func(int, int) int
	switch op {
	case "add":
		operation = func(a, b int) int {
			return a + b
		}
		break
	case "subtract":
		operation = func(a, b int) int {
			return a - b
		}
		break
	case "multiply":
		operation = func(a, b int) int {
			return a * b
		}
		break
	case "divide":
		operation = func(a, b int) int {
			if b != 0 {
				return a / b
			}
			return 0
		}
		break
	default:
		return nil
	}
	return &operation
}

// func main() {
// 	add := GetMathOperation("add")
// 	if add != nil {
// 		result := (*add)(2, 3)
// 		fmt.Println(result) // 5
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	subtract := GetMathOperation("subtract")
// 	if subtract != nil {
// 		result := (*subtract)(5, 2)
// 		fmt.Println(result) // 3
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	multiply := GetMathOperation("multiply")
// 	if multiply != nil {
// 		result := (*multiply)(3, 4)
// 		fmt.Println(result) // 12
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	divide := GetMathOperation("divide")
// 	if divide != nil {
// 		result := (*divide)(10, 2)
// 		fmt.Println(result) // 5
// 		result = (*divide)(10, 0)
// 		fmt.Println(result) // 0
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}

// 	invalid := GetMathOperation("mod")
// 	if invalid != nil {
// 		result := (*invalid)(10, 3)
// 		fmt.Println(result)
// 	} else {
// 		fmt.Println("Invalid operation")
// 	}
// }
