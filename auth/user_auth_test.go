package auth

import "testing"

var ua = NewUserAuth("../auth.txt")

func TestMustGetOriginAuth(t *testing.T) {
	t.Log(string(ua.MustGetOriginAuth()))
}

func TestMustGetHashAuth(t *testing.T) {
	t.Log(string(ua.MustGetHashAuth()))
}

func TestMustIsAuth(t *testing.T) {
	flag := ua.MustIsAuth([]byte("$2a$10$J7/xUvxCIhXweT74O6pYPuDGUvZFy3lfFDpLxOPnTeuPUOJMFzfXa"))

	if !flag {
		t.Fatal()
	}
}
