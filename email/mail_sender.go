package email

import (
	"encoding/json"
	"log"
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
	mail "gopkg.in/gomail.v2"
)

type EmailProvider struct {
	Sender string
	Token  string
	Host   string
}

func mustGetProvider(uri string) EmailProvider {
	data, err := os.ReadFile(uri)

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

type MailSender struct {
	EmailProvider
}

func NewMailSender(providerUri string) MailSender {
	p := mustGetProvider(providerUri)
	return MailSender{
		EmailProvider: p,
	}
}

type Email struct {
	Receiver string
	Subject  string
	HTMLBody string
}

func (sender MailSender) SendMail(em Email) error {
	provider := sender.EmailProvider

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
