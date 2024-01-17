package controller

import (
	"github.com/gin-gonic/gin"
	. "goregister.com/app/jwt"
)

type JwtController struct {
	service JwtService
	group   *gin.RouterGroup
}

func NewJwtController(service JwtService, router *gin.Engine) JwtController {
	return JwtController{
		service: service,
		group:   router.Group("api/v1/jwt/"),
	}
}

func (ctr JwtController) Run() {
	ctr.Login()
}

func (ctr JwtController) Login() {
	ctr.group.POST("login/", ctr.service.Login)
}
