package jwt

import (
	"testing"

	. "goregister.com/app/data"
)

func TestJwtProvider(t *testing.T) {
	var provider JwtProvider
	KEY := "aaa"
	jwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTE5NDEwODcsImlkIjoiMTIzIn0.W0-80nUDBOj_AkHKSJMeFH3a8raaC5XIp_uBKfz2L-g"
	t.Run("test must get jwt", func(t *testing.T) {
		user := User{
			Id: "123",
		}

		token, err := provider.GetJWT([]byte(KEY), user.Id)
		if err == nil {
			t.Log(token)
		} else {
			t.Log(err)
		}
	})
	t.Run("test must is jwt verified", func(t *testing.T) {
		res := provider.MustIsJWTVerified([]byte(KEY), jwtToken)
		if !res {
			t.Fatal("jwt not verified")
		}
	})
	t.Run("test parse jwt", func(t *testing.T) {
		_, err := parseJWT([]byte(KEY), jwtToken)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test must get jwt claims", func(t *testing.T) {
		claims := provider.MustGetJWTClaims([]byte(KEY), jwtToken)
		if claims == nil {
			t.Fatal()
		} else {
			t.Log(claims)
			t.Log("id: ", claims["id"])
		}
	})
}
