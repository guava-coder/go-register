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
	userRepo UserRepository
}

func NewEmailService(userRepo UserRepository) EmailService {
	return EmailService{
		userRepo: userRepo,
	}
}

func getVerificationMailForm(receiver User) Email {
	code := receiver.Auth
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

func sendVerificationMail(receiver User, ctx *gin.Context) {
	var sender MailSender
	err := sender.SendMail(getVerificationMailForm(receiver))
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

func CheckUserExist(user User, notExist func(), exist func(User)) {
	if user.Email == "" {
		notExist()
	} else {
		exist(user)
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
	var handler EmailHandler

	addFakeUserAuthAndSendMail := func(ctx *gin.Context, user User) {
		CheckUserExist(serv.userRepo.QueryById(user.Id),
			func() {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"Response": "User email incorrect.",
				})
			},
			func(u User) {
				user := serv.userRepo.UpdateUserAuth(User{
					Id:   u.Id,
					Auth: RandStringBytes(6),
				})
				sendVerificationMail(user, ctx)
			},
		)
	}

	handler.VerifyUserEmail(ctx, addFakeUserAuthAndSendMail)
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
