package email

import (
	"testing"

	. "goregister.com/app/data"
)

func TestService(t *testing.T) {
	t.Run("test get verification mail form", func(t *testing.T) {
		form := getVerificationMailForm(User{
			Email: "mark@mail.com",
			Auth:  "123",
		})
		if form.Subject == "" {
			t.Fatal()
		}
	})
	t.Run("test random string bytes", func(t *testing.T) {
		v := string(RandStringBytes(4, "1234567890"))
		if v == "" {
			t.Fatal()
		}
		t.Log(v)
	})
}
