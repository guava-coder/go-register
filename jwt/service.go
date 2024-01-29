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
	userAuth UserAuth
}

func NewJwtService(userRepo UserRepository, userAuth UserAuth) JwtService {
	return JwtService{
		userRepo: userRepo,
		userAuth: userAuth,
	}
}

func (serv JwtService) readAndHandleRequestBody(ctx *gin.Context, op func(User)) {
	ReadAndHandleRequestBody[User](ctx, op)
}

func (serv JwtService) Login(ctx *gin.Context) {
	getJWT := func(key []byte, user User) {
		var provider JwtProvider
		token, err := provider.GetJWT(key, user.Id)
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
			if serv.userAuth.MustIsAuth([]byte(user.Auth)) {
				getJWT(serv.userAuth.MustGetOriginAuth(), user)
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"Response": "User doesn't have functional authorized key",
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
				"Response": "Not a User",
			})
		} else {
			verifyUser(us, user)
		}
	}
	serv.readAndHandleRequestBody(ctx, getToken)
}

func (serv JwtService) VerifyBearerToken(ctx *gin.Context) {
	verf := NewBearerVerfier(serv.userAuth, ctx)
	verf.VerifyBearerToken(
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Authorized",
			})
		},
	)
}
