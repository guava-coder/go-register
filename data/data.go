package data

type Data interface {
	User
}

type User struct {
	Id       string
	Name     string
	Email    string
	Bio      string
	Password string
	Auth     string
	TempCode string
}
