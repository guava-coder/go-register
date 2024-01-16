package auth

import (
	"embed"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//go:embed auth.txt
var embedKey embed.FS

func (ua UserAuth) MustGetOriginAuth() []byte {
	key, err := embedKey.ReadFile("auth.txt")
	if err != nil {
		log.Println(err)
	}
	return key
}

type UserAuth struct{}

func (ua UserAuth) MustGetHashAuth() []byte {
	bytes, err := bcrypt.GenerateFromPassword(ua.MustGetOriginAuth(), 0)

	if err != nil {
		log.Println(err)
	}
	return bytes
}

func (ua UserAuth) MustIsAuth(auth []byte) bool {
	err := bcrypt.CompareHashAndPassword(auth, ua.MustGetOriginAuth())
	if err == nil {
		return true
	} else {
		return false
	}
}
