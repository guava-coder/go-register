package controller

import (
	"net/http"
	"strings"
	"testing"
)

func TestJwtController(t *testing.T) {
	prefix := "http://127.0.0.1:8080/api/v1/jwt/"
	t.Run("test login", func(t *testing.T) {
		body := `{"Email": "mark@mail.com", "Password": "123"}`
		req, err := http.NewRequest(http.MethodPost, prefix+"login/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		HandleTestRequest(req, t)
	})
}
