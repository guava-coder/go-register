package email

import (
	"testing"

	. "goregister.com/app/data"
)

func TestGetVerificationMailForm(t *testing.T) {
	form := getVerificationMailForm(User{
		Email: "mark@mail.com",
		Auth:  "123",
	})
	t.Log(form)
}

func TestRandStringBytes(t *testing.T) {
	t.Log(RandStringBytes(6))
}

func TestCheckUserExist(t *testing.T) {
	CheckUserExist((User{
		Email: "lisa@mail.com",
		Auth:  "6666",
	}),
		func() {
			t.Fatal("User not exist")
		},
		func(u User) {
			u.Auth = RandStringBytes(6)
			form := getVerificationMailForm(u)
			t.Log(form)
		},
	)
}
