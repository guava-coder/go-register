package controller

import (
	"net/http"
	"strings"
	"testing"
)

func TestEmailController(t *testing.T) {
	prefix := "http://127.0.0.1:8080/api/v1/email/"
	t.Run("test send verification email", func(t *testing.T) {
		body := `{"Id":"a01","Email":"mark@mail.com"}`
		req, err := http.NewRequest(http.MethodPost, prefix+"send/verification/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		HandleTestRequest(req, t)
	})
	t.Run("test verify email", func(t *testing.T) {
		body := `{"Email": "mark@mail.com"}`
		req, err := http.NewRequest(http.MethodPost, prefix+"verify/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTE5NDQ1NDMsImlkIjoiYTAxIn0.hx_ID4yn58s4d7QdK3jXcbAL-U6YhjTIEOo5qokfeyQ")
		HandleTestRequest(req, t)
	})
}
