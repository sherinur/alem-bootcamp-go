package firststruct

type User struct {
	Name     string
	password string
}

func NewUser(name, password string) User {
	return User{
		Name:     name,
		password: password,
	}
}

func (u *User) GetPassword() string {
	return u.password
}
