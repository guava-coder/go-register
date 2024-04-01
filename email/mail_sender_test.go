package email

import "testing"

func TestMailSender(t *testing.T) {
	recevier := "###@gmail.com"
	providerUri := "../provider.json"
	sender := NewMailSender(providerUri)
	t.Run("test send mail", func(t *testing.T) {
		em := Email{
			Receiver: recevier,
			Subject:  "Go mail test",
			HTMLBody: "Go mail test, don't reply",
		}
		err := sender.SendMail(em)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test verify email", func(t *testing.T) {
		res := sender.VerifyEmail(recevier)
		if res == nil {
			t.Fatal()
		}
	})
	t.Run("test must get provider", func(t *testing.T) {
		provider := mustGetProvider(providerUri)
		if provider.Sender == "" {
			t.Fatal()
		}
	})
}
