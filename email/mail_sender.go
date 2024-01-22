package email

import (
	"embed"
	"encoding/json"
	"log"

	emailverifier "github.com/AfterShip/email-verifier"
	mail "gopkg.in/gomail.v2"
)

type EmailProvider struct {
	Sender string
	Token  string
	Host   string
}

//go:embed provider.json
var embedKey embed.FS

func mustGetProvider() EmailProvider {
	data, err := embedKey.ReadFile("provider.json")
	if err != nil {
		log.Println(err)
	}
	var provider EmailProvider
	err = json.Unmarshal(data, &provider)
	if err != nil {
		log.Println(err)
	}
	return provider
}

type MailSender struct{}

type Email struct {
	Receiver string
	Subject  string
	HTMLBody string
}

func (sender MailSender) SendMail(em Email) error {
	provider := mustGetProvider()

	m := mail.NewMessage()
	m.SetHeader("From", provider.Sender)

	m.SetHeader("To", em.Receiver)
	m.SetHeader("Subject", em.Subject)

	m.SetBody("text/html", em.HTMLBody)

	d := mail.NewDialer(provider.Host, 587, provider.Sender, provider.Token)

	err := d.DialAndSend(m)
	return err
}

func (sender MailSender) VerifyEmail(email string) *emailverifier.Result {
	verifier := emailverifier.NewVerifier()

	ret, err := verifier.Verify(email)
	if err != nil {
		log.Println("ERROR: " + err.Error())
	}
	return ret
}
