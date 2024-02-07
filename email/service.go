package email

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	. "goregister.com/app/data"
	. "goregister.com/app/user"
)

type EmailService struct {
	userRepo   UserRepository
	mailSender MailSender
}

func NewEmailService(userRepo UserRepository, m MailSender) EmailService {
	return EmailService{
		userRepo:   userRepo,
		mailSender: m,
	}
}

func getVerificationMailForm(receiver User) Email {
	code := receiver.TempCode
	body := fmt.Sprintf(`
				<h3>Go-Register email verification</h3>
				<p>Using the code below to verify your email. 
				Don't reply this email.</p>
				<h1> %v </h1>
			`, code)
	form := Email{
		Subject:  "Go-Register email Verification",
		HTMLBody: body,
	}

	form.Receiver = receiver.Email
	return form
}

func (serv EmailService) sendVerificationMail(receiver User, ctx *gin.Context) {
	err := serv.mailSender.SendMail(getVerificationMailForm(receiver))
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

func RandStringBytes(length int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (serv EmailService) SendVerificationEmail(ctx *gin.Context) {
	addTempCodeAndSendMail := func(id string) {
		updatedU, err := serv.userRepo.UpdateTempCode(User{
			Id:       id,
			TempCode: RandStringBytes(6),
		})
		if err == nil {
			serv.sendVerificationMail(updatedU, ctx)
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "System failed to generate temporary code",
			})
		}
	}
	SendMailToUser := func(ctx *gin.Context, user User) {
		foundU, err := serv.userRepo.QueryById(user.Id)
		if err == nil {
			if foundU.Email == user.Email {
				addTempCodeAndSendMail(foundU.Id)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": "User email address invalid",
				})
			}
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Response": err.Error(),
			})
		}
	}

	var handler EmailHandler
	handler.VerifyUserEmail(ctx, SendMailToUser)
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
