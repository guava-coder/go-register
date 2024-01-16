package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	. "goregister.com/app/auth"
	. "goregister.com/app/data"
	. "goregister.com/app/request"
	. "goregister.com/app/user"
)

type JwtService struct {
	userRepo UserRepository
}

func NewJwtService(userRepo UserRepository) JwtService {
	return JwtService{
		userRepo: userRepo,
	}
}

func (serv JwtService) readAndHandleRequestBody(ctx *gin.Context, op func(User)) {
	ReadAndHandleRequestBody[User](ctx, op)
}

func (serv JwtService) Login(ctx *gin.Context) {
	getJWT := func(key []byte, user User) {
		token, err := GetJWT(key, user.Id)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{"Token": token})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "ERROR: " + err.Error(),
			})
		}
	}

	verifyUser := func(input User, user User) {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if err == nil {
			var ua UserAuth
			if ua.MustIsAuth([]byte(user.Auth)) {
				getJWT(ua.MustGetOriginAuth(), user)
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"Response": "User doesn't have an authorized key " + user.Auth,
				})
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Response": "Password incorrect",
			})
		}
	}

	getToken := func(us User) {
		user := serv.userRepo.QueryByInfo(us)
		if user.Id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Unauthorized User",
			})
		} else {
			verifyUser(us, user)
		}
	}
	serv.readAndHandleRequestBody(ctx, getToken)
}
