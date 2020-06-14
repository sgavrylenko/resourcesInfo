package users

type User struct {
	ID       int
	Username string
}

func Login(user, password string) []User {
	var users []User
	users = append(users, User{ID: 0, Username: "test"})
	return users
}
