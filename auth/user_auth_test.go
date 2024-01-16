package auth

import "testing"

var ua UserAuth

func TestMustGetOriginAuth(t *testing.T) {
	t.Log(string(ua.MustGetOriginAuth()))
}

func TestMustGetHashAuth(t *testing.T) {
	t.Log(string(ua.MustGetHashAuth()))
}

func TestMustIsAuth(t *testing.T) {
	flag := ua.MustIsAuth([]byte("$2a$10$PcbcpTWv/BaI5cvCcATzV.bc0zibIN5Q6S5xOpl0eRLbjDIaAfyZq"))

	if !flag {
		t.Fatal()
	}
}
