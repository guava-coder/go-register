package controller

import (
	"github.com/gin-gonic/gin"
	email "goregister.com/app/email"
	jwt "goregister.com/app/jwt"
	user "goregister.com/app/user"
)

type UserController struct {
	service user.UserService
	group   *gin.RouterGroup
}

func NewUserController(service user.UserService, router *gin.Engine) UserController {
	return UserController{
		service: service,
		group:   router.Group("api/v1/user/"),
	}
}

func (ctr UserController) Run() {
	ctr.addUser()
	ctr.queryById()
	ctr.updateUserAuth()
	ctr.updatePassword()
	ctr.updateUserInfo()

	ctr.checkPassword()
}

func (ctr UserController) addUser() {
	ctr.group.POST("add/", func(ctx *gin.Context) {
		var handler email.EmailHandler
		handler.VerifyUserEmail(ctx, ctr.service.AddUser)
	})
}

func (ctr UserController) queryById() {
	ctr.group.POST("query/", func(ctx *gin.Context) {
		verifier := jwt.NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.QueryById)
	})
}

func (ctr UserController) updateUserAuth() {
	ctr.group.PUT("auth/", ctr.service.UpdateUserAuth)
}

func (ctr UserController) updatePassword() {
	ctr.group.PUT("password/", func(ctx *gin.Context) {
		verifier := jwt.NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.UpdatePassword)
	})
}

func (ctr UserController) updateUserInfo() {
	ctr.group.PUT("update/", func(ctx *gin.Context) {
		verifier := jwt.NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.UpdateUserInfo)
	})
}

func (ctr UserController) checkPassword() {
	ctr.group.POST("check/password/", func(ctx *gin.Context) {
		verifier := jwt.NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.CheckPassword)
	})
}
