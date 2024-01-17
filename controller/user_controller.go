package controller

import (
	"github.com/gin-gonic/gin"
	. "goregister.com/app/user"
)

type UserController struct {
	service UserService
	group   *gin.RouterGroup
}

func NewUserController(service UserService, router *gin.Engine) UserController {
	return UserController{
		service: service,
		group:   router.Group("api/v1/user"),
	}
}

func (ctr UserController) Run() {
	ctr.AddUser()
	ctr.QueryById()
	ctr.UpdateUserAuth()
}

func (ctr UserController) AddUser() {
	ctr.group.POST("/add", ctr.service.AddUser)
}

func (ctr UserController) QueryById() {
	ctr.group.POST("/query", ctr.service.QueryById)
}

func (ctr UserController) UpdateUserAuth() {
	ctr.group.POST("/auth", ctr.service.UpdateUserAuth)
}
