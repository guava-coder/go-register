package controller

import (
	"github.com/gin-gonic/gin"
	. "goregister.com/app/email"
)

type EmailController struct {
	service EmailService
	group   *gin.RouterGroup
}

func NewEmailController(service EmailService, router *gin.Engine) EmailController {
	return EmailController{
		service: service,
		group:   router.Group("/api/v1/email"),
	}
}

func (ctr EmailController) Run() {
	ctr.SendVerificationEmail()
	ctr.VerifyEmail()
}

func (ctr EmailController) SendVerificationEmail() {
	ctr.group.POST(
		"/send",
		ctr.service.SendVerificationEmail,
	)
}

func (ctr EmailController) VerifyEmail() {
	ctr.group.POST("/verify", ctr.service.VerifyEmail)
}
