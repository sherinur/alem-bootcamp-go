package bootcamp

import (
	"bootcamp/firststruct"
)

type Cart struct {
	Owner *firststruct.User
	Items []string
}

func CreateMyCart(owner *firststruct.User, items []string) *Cart {
	return &Cart{
		owner,
		items,
	}
}

// func main() {
// 	user := firststruct.NewUser("Alice", "securePassword123")
// 	items := []string{"item1", "item2", "item3"}

// 	cart := CreateMyCart(&user, items)
// 	fmt.Println("Cart owner:", cart.Owner.Name)
// 	// Cart owner: Alice
// 	fmt.Println("Cart items:", cart.Items)
// 	// Cart items: [item1 item2 item3]
// }
