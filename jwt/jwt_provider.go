package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtProvider struct{}

func (pro JwtProvider) GetJWT(key []byte, id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["authorized"] = true
	claims["id"] = id

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

func parseJWT(key []byte, token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return key, nil
	})
	return t, err
}

func (pro JwtProvider) MustIsJWTVerified(key []byte, token string) bool {
	t, err := parseJWT(key, token)
	if err != nil {
		log.Println(err)
		return false
	}
	if t.Valid {
		return true
	} else {
		log.Println("Invalid token")
		return false
	}
}

func (pro JwtProvider) MustGetJWTClaims(key []byte, token string) jwt.MapClaims {
	res, err := parseJWT([]byte(key), token)

	if err == nil {
		claims, ok := res.Claims.(jwt.MapClaims)
		if ok && res.Valid {
			return claims
		}
	}
	return nil
}
