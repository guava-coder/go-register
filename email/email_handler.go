package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "goregister.com/app/data"
	. "goregister.com/app/request"
)

func readAndHandleRequestBody(ctx *gin.Context, operation func(User)) {
	ReadAndHandleRequestBody[User](ctx, operation)
}

type EmailHandler struct{}

func (handler EmailHandler) VerifyUserEmail(ctx *gin.Context, authOp func(*gin.Context, User)) {
	readAndHandleRequestBody(ctx, func(u User) {
		var sender MailSender
		res := sender.VerifyEmail(u.Email)

		if res != nil && res.Syntax.Valid {
			authOp(ctx, u)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Response": "Email invaild.",
			})
		}
	})
}
