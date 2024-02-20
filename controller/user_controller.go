package controller

import (
	"github.com/gin-gonic/gin"
	. "goregister.com/app/email"
	. "goregister.com/app/jwt"
	. "goregister.com/app/user"
)

type UserController struct {
	service UserService
	group   *gin.RouterGroup
}

func NewUserController(service UserService, router *gin.Engine) UserController {
	return UserController{
		service: service,
		group:   router.Group("api/v1/user/"),
	}
}

func (ctr UserController) Run() {
	ctr.AddUser()
	ctr.QueryById()
	ctr.UpdateUserAuth()
	ctr.UpdatePassword()
	ctr.UpdateUserInfo()
}

func (ctr UserController) AddUser() {
	ctr.group.POST("add/", func(ctx *gin.Context) {
		var handler EmailHandler
		handler.VerifyUserEmail(ctx, ctr.service.AddUser)
	})
}

func (ctr UserController) QueryById() {
	ctr.group.POST("query/", func(ctx *gin.Context) {
		verifier := NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.QueryById)
	})
}

func (ctr UserController) UpdateUserAuth() {
	ctr.group.PUT("auth/", ctr.service.UpdateUserAuth)
}

func (ctr UserController) UpdatePassword() {
	ctr.group.PUT("password/", func(ctx *gin.Context) {
		verifier := NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.UpdatePassword)
	})
}

func (ctr UserController) UpdateUserInfo() {
	ctr.group.POST("update/", func(ctx *gin.Context) {
		verifier := NewBearerVerfier(ctr.service.UserAuth, ctx)
		verifier.ExtractUserIdFromBearer(ctr.service.UpdateUserInfo)
	})
}
