package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	auth "goregister.com/app/auth"
	data "goregister.com/app/data"
	req "goregister.com/app/request"
	user "goregister.com/app/user"
)

type JwtService struct {
	userRepo user.UserRepository
	userAuth auth.UserAuth
}

func NewJwtService(userRepo user.UserRepository, userAuth auth.UserAuth) JwtService {
	return JwtService{
		userRepo: userRepo,
		userAuth: userAuth,
	}
}

func (serv JwtService) readAndHandleRequestBody(ctx *gin.Context, op func(data.User)) {
	req.ReadAndHandleRequestBody(ctx, op)
}

// Login performs the User login functionality.
//
// ctx *gin.Context
func (serv JwtService) Login(ctx *gin.Context) {
	getJWT := func(key []byte, user data.User) {
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

	checkUserAuthorized := func(user data.User) {
		if user.Auth != "" && serv.userAuth.MustIsAuth([]byte(user.Auth)) {
			getJWT(serv.userAuth.MustGetOriginAuth(), user)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Response": "User doesn't have functional authorized key",
				"Id":       user.Id,
			})
		}
	}

	comparePassword := func(user data.User, input data.User) {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
		if err == nil {
			checkUserAuthorized(user)
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Response": "Password incorrect",
			})
		}
	}

	getToken := func(input data.User) {
		user, err := serv.userRepo.QueryByInfo(input)
		if err == nil {
			comparePassword(user, input)
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "User not found",
			})
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
