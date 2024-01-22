package email

import "testing"

func TestMustReadEmailForm(t *testing.T) {
	t.Log(MustReadEmailForm())
}
