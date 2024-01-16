package jwt

import (
	"testing"

	. "goregister.com/app/data"
)

var (
	KEY       = "aaa"
	JWT_TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MDUzOTQ4NjgsImlkIjoiMTIzIn0.Dexfssj9DMOeMF2tUwTkU0H2NGDuELNkSRHXE7iHX6c"
)

func TestMustGetGetJWT(t *testing.T) {
	user := User{
		Id: "123",
	}

	token, err := GetJWT([]byte(KEY), user.Id)
	if err == nil {
		t.Log(token)
	} else {
		t.Log(err)
	}
}

func TestVerifyJWTToken(t *testing.T) {
	res := MustVerifyJWT([]byte(KEY), JWT_TOKEN)
	if res == false {
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
	claims := MustGetJWTClaims([]byte(KEY), JWT_TOKEN)
	if claims == nil {
		t.Fatal()
	} else {
		t.Log(claims)
		t.Log("id: ", claims["id"])
	}
}
