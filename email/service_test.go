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
	t.Log(string(RandStringBytes(4, "1234567890")))
}
