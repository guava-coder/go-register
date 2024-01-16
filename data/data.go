package data

type Data interface {
	User
}

type User struct {
	Id       string
	Name     string
	Email    string
	Phone    string
	BirthDay string
	Gender   string
	Password string
	Auth     string
}
