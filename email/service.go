package email

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Response": "Sending a verification mail failed, please try again.",
		})
	}
}

const (
	UpperCaseNums = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	EngCharacters = "ABCDEFGHIJKLMNOPQRSTUVWXYZqwertyuiopasdfghjklzxcvbnm0123456789"
)

func RandStringBytes(length int, sample string) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = sample[rand.Intn(len(sample))]
	}
	return b
}

func (serv EmailService) SendVerificationEmail(ctx *gin.Context) {
	addTempCodeAndSendMail := func(id string) {
		updatedU, err := serv.userRepo.UpdateTempCode(User{
			Id:       id,
			TempCode: string(RandStringBytes(6, UpperCaseNums)),
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

func NewPasswordMail(receiver string, password string) Email {
	body := fmt.Sprintf(`
				<h3>Go-Register Recover password</h3>
				<p>We send you a new password, use it to login. 
				Don't reply this email.</p>
				<h1> %v </h1>
			`, password)
	form := Email{
		Subject:  "Go-Register Recover Password",
		HTMLBody: body,
	}

	form.Receiver = receiver
	return form
}

func (serv EmailService) SendTemporaryPassword(ctx *gin.Context) {
	sendMailToUser := func(u User, password string) {
		err := serv.mailSender.SendMail(NewPasswordMail(u.Email, password))
		if err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"Response": "We sended a new password for you, please check your email and use it to login.",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Response": "Sending mail failed, please try again.",
			})
		}
	}

	checkTempCode := func(ctx *gin.Context, u User) {
		if serv.userRepo.IsTempCodeCorrect(u) {
			unhash := RandStringBytes(8, EngCharacters)

			newpsw, err := bcrypt.GenerateFromPassword(unhash, 0)
			u.Password = string(newpsw)

			res, err := serv.userRepo.UpdatePassword(u)

			if err == nil {
				sendMailToUser(res, string(unhash))
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": err.Error(),
				})
			}
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{
				"Response": "code incorrect",
			})
		}
	}

	var handler EmailHandler
	handler.VerifyUserEmail(ctx, checkTempCode)
}
