package controller

import (
	"github.com/gin-gonic/gin"
	email "goregister.com/app/email"
)

type EmailController struct {
	service email.EmailService
	group   *gin.RouterGroup
}

func NewEmailController(service email.EmailService, router *gin.Engine) EmailController {
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
		"/send/verification",
		ctr.service.SendVerificationEmail,
	)
}

func (ctr EmailController) VerifyEmail() {
	ctr.group.POST("/verify", ctr.service.VerifyEmail)
}
