package controller

import (
	"github.com/gin-gonic/gin"
	jwt "goregister.com/app/jwt"
)

type JwtController struct {
	service jwt.JwtService
	group   *gin.RouterGroup
}

func NewJwtController(service jwt.JwtService, router *gin.Engine) JwtController {
	return JwtController{
		service: service,
		group:   router.Group("api/v1/jwt/"),
	}
}

func (ctr JwtController) Run() {
	ctr.Login()
	ctr.VerifyBearerToken()
}

func (ctr JwtController) Login() {
	ctr.group.POST("login/", ctr.service.Login)
}

func (ctr JwtController) VerifyBearerToken() {
	ctr.group.POST("verify/", ctr.service.VerifyBearerToken)
}
