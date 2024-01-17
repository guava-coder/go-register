package jwt

import (
	"testing"

	. "goregister.com/app/data"
)

var (
	provider  JwtProvider
	KEY       = "aaa"
	JWT_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDU0NzQ2OTQsImlkIjoiMTIzIn0.9ft0sQZoZFxcW4i0wE8oUbS5fzy215uyyXOBDPJno3g"
)

func TestMustGetGetJWT(t *testing.T) {
	user := User{
		Id: "123",
	}

	token, err := provider.GetJWT([]byte(KEY), user.Id)
	if err == nil {
		t.Log(token)
	} else {
		t.Log(err)
	}
}

func TestVerifyJWTToken(t *testing.T) {
	res := provider.MustIsJWTVerified([]byte(KEY), JWT_TOKEN)
	if !res {
		t.Fatal()
	}
}

func TestParseJWT(t *testing.T) {
	res, err := parseJWT([]byte(KEY), JWT_TOKEN)
	if err == nil {
		t.Log(res)
	} else {
		t.Fatal(err)
	}
}

func TestMustGetJWTClaims(t *testing.T) {
	claims := provider.MustGetJWTClaims([]byte(KEY), JWT_TOKEN)
	if claims == nil {
		t.Fatal()
	} else {
		t.Log(claims)
		t.Log("id: ", claims["id"])
	}
}
