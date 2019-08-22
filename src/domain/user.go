package domain

type User struct{
	Name string
	Username string
	Email string
	Password string
}

func NewUser(name,username,email,password string) *User{
	user := User{
		name,
		username,
		email,
		password,
	}
	return &user
}