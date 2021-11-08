package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestNew(t *testing.T) {
	a.Initialize(
		goDotEnvVariable("DB_USER"),
		goDotEnvVariable("DB_PASSWORD"),
		goDotEnvVariable("DB_HOST"),
		goDotEnvVariable("DB_NAME"),
	)
}

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
	a.DB.Exec("ALTER SEQUENCE todo_id_seq RESTART WITH 1")
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

func TestAddTodoList(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"Todolist":"test product"}`)
	req, _ := http.NewRequest("POST", "/todolist", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["todolist"] != "test product" {
		t.Errorf("Expected todolist to be 'test product'. Got '%v'", m["todolist"])
	}

	// the id is compared to 1.0 because JSON unmarshaling converts numbers to
	// floats, when the target is a map[string]interface{}
	if m["id"] != 1.0 {
		t.Errorf("Expected todo ID to be '1'. Got '%v'", m["id"])
	}
}

func TestGetTodoListitems(t *testing.T) {
	clearTable()
	addtodoitemsintable(1)

	req, _ := http.NewRequest("GET", "/todolist", nil)
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	bytes := []byte(response.Body.String())

	// Unmarshal string into structs.
	var m1 []todo
	json.Unmarshal(bytes, &m1)
	if m1[0].ID != 1.0 {
		t.Errorf("Expected todo ID to be '1'. Got '%v'", m1[0].ID)
	}

}

func addtodoitemsintable(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO todo(Todolist) VALUES($1)", "Product")
	}
}
