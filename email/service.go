package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "goregister.com/app/data"
	. "goregister.com/app/user"
)

type EmailService struct {
	userRepo UserRepository
}

func NewEmailService(userRepo UserRepository) EmailService {
	return EmailService{
		userRepo: userRepo,
	}
}

func (serv EmailService) SendVerificationEmail(ctx *gin.Context) {
	sendVerificationMail := func(receiver string) {
		form := Email{
			Subject:  "Go-Register email Verification",
			HTMLBody: string(MustReadEmailForm()),
		}

		form.Receiver = receiver

		var sender MailSender
		err := sender.SendMail(form)
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "A verification mail has sended, please check your email and verify it.",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": "Sending a verification mail failed, please try again.",
			})
		}
	}

	var handler EmailHandler
	handler.VerifyUserEmail(ctx,
		func(ctx *gin.Context, user User) {
			u := serv.userRepo.QueryByInfo(user)
			if u.Email == "" {
				sendVerificationMail(user.Email)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": "User email incorrect.",
				})
			}
		},
	)
}

func (serv EmailService) VerifyEmail(ctx *gin.Context) {
	var handler EmailHandler
	handler.VerifyUserEmail(ctx,
		func(ctx *gin.Context, user User) {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "Email verified.",
			})
		},
	)
}
