package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestNew(t *testing.T) {
	a.Initialize(
		"postgres",
		"postgres",
		"testing")
}

func TestFail(t *testing.T) {
	t.Errorf("Expected error")

}
func Test_testAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"data":"test"}` {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}
