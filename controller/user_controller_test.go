package controller

import (
	"net/http"
	"testing"
)

func TestUserController(t *testing.T) {
	prefix := "http://127.0.0.1:8080/api/v1/user/"
	t.Run("test query by id", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, prefix+"query/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTE5NDQ1NDMsImlkIjoiYTAxIn0.hx_ID4yn58s4d7QdK3jXcbAL-U6YhjTIEOo5qokfeyQ")
		HandleTestRequest(req, t)
	})
}
