package email

import "testing"

var (
	recevier    = "###@gmail.com"
	providerUri = "../provider.json"
	sender      = NewMailSender(providerUri)
)

func TestMustGetProvider(t *testing.T) {
	provider := mustGetProvider(providerUri)
	if provider.Sender == "" {
		t.Fatal()
	} else {
		t.Log(provider)
	}
}

func TestSendMail(t *testing.T) {
	em := Email{
		Receiver: recevier,
		Subject:  "Go mail test",
		HTMLBody: "Go mail test, don't reply",
	}
	err := sender.SendMail(em)
	if err != nil {
		t.Fatal(err)
	}
}

func TestVerifyEmail(t *testing.T) {
	res := sender.VerifyEmail(recevier)
	if res == nil {
		t.Fatal()
	}
}
