package bootcamp

import (
	"bootcamp/firststruct"

	"github.com/alem-platform/ap"
)

// func main() {
// 	userAlem := firststruct.NewUser("Alem", "hello.alem")
// 	PrintUserInfo(userAlem)
// 	// Output:
// 	// Name: Alem
// 	// HasPassword: yes

// 	userDias := firststruct.NewUser("Dias", "")
// 	PrintUserInfo(userDias)
// 	// Output:
// 	// Name: Dias
// 	// HasPassword: no
// }

func outputString(s string) {
	for _, c := range s {
		ap.PutRune(c)
	}
}

func PrintUserInfo(u firststruct.User) {
	hasPassword := !(u.GetPassword() == "")
	outputString("Name: " + u.Name + "\n")
	outputString("HasPassword: ")

	if hasPassword {
		outputString("yes\n")
	} else {
		outputString("no\n")
	}
}
