package auth

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type UserAuth struct {
	keyUri string
}

func NewUserAuth(keyUri string) UserAuth {
	return UserAuth{
		keyUri: keyUri,
	}
}

func (ua UserAuth) MustGetOriginAuth() []byte {
	key, err := os.ReadFile(ua.keyUri)
	if err != nil {
		log.Println(err)
	}
	return key
}

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
