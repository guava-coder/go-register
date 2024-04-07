package controller

import (
	"io"
	"net/http"
	"testing"
)

func assertStatusOk(res *http.Response, t *testing.T) {
	if res.StatusCode != 200 {
		t.Fatal("status code should be 200")
	}
}

func logResponseBody(res *http.Response, t *testing.T) {
	value, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(value))
}

func getTestResponse(t *testing.T, req *http.Request) (res *http.Response) {
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

func HandleTestRequest(req *http.Request, t *testing.T) {
	res := getTestResponse(t, req)
	HandleTestResponse(res, t)
}

func HandleTestResponse(res *http.Response, t *testing.T) {
	defer res.Body.Close()
	logResponseBody(res, t)
	assertStatusOk(res, t)
}
