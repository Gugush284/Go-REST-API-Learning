package model_user

// User struct ...
type User struct {
	ID       int
	Login    string
	Password string
}

func New() *User {
	return &User{}
}
