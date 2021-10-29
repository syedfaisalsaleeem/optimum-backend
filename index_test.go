package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestNew(t *testing.T) {
	a.Initialize(
		"postgres",
		"faisal",
		"postgres")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS todo
(
    id SERIAL,
    Todolist TEXT
)`

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func Test_testAPI(t *testing.T) {
	req, _ := http.NewRequest("GET", "/test", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `[]` {
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

func TestDBConnection(t *testing.T) {
	ensureTableExists()
}

func clearTable() {
	a.DB.Exec("DELETE FROM todo")
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/todolist", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}
