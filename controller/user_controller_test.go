package controller

import (
	"net/http"
	"strings"
	"testing"
)

var prefix = "http://127.0.0.1:8080/api/v1/user/"

func TestUserController(t *testing.T) {
	bearer := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTE5NDQ1NDMsImlkIjoiYTAxIn0.hx_ID4yn58s4d7QdK3jXcbAL-U6YhjTIEOo5qokfeyQ"
	t.Run("test query by id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, prefix+"query/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", bearer)
		HandleTestRequest(req, t)
	})

	t.Run("test update user info", func(t *testing.T) {
		body := `{"Name":"Dora","Bio":"Test User"}`
		req, err := http.NewRequest(http.MethodPut, prefix+"update/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", bearer)
		HandleTestRequest(req, t)
	})
	t.Run("test check password", func(t *testing.T) {
		body := `{"Password":"123"}`
		req, err := http.NewRequest(http.MethodPost, prefix+"check/password/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", bearer)
		HandleTestRequest(req, t)
	})
}

func TestUserControllerModifieUser(t *testing.T) {
	t.Run("test add user", func(t *testing.T) {
		body := `{"Id":"a99","Name":"Dora","Email":"dora@mail.com","Bio":"Test User","Password":"1234","Auth":"none"}`
		req, err := http.NewRequest(http.MethodPost, prefix+"add/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		HandleTestRequest(req, t)
	})
	t.Run("test update user auth", func(t *testing.T) {
		body := `{"Id": "69b52219-7886-4163-bf10-740d33a0ebc7","TempCode":"WMZT1Y"}`
		req, err := http.NewRequest(http.MethodPut, prefix+"auth/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		HandleTestRequest(req, t)
	})
	t.Run("test update password", func(t *testing.T) {
		body := `{"Id": "69b52219-7886-4163-bf10-740d33a0ebc7","Password":"1234"}`
		req, err := http.NewRequest(http.MethodPut, prefix+"password/", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		HandleTestRequest(req, t)
	})
}
