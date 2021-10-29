package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}
type test_data struct {
	Data string `json:"data"`
}

func (a *App) Initialize(user, password, dbname string) {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	http.ListenAndServe(":8010", a.Router)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	data1 := test_data{
		Data: "test",
	}
	respondWithJSON(w, http.StatusOK, data1)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/test", a.getProducts).Methods("GET")
}

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"postgres",
		"testing")

	a.Run(":8010")
}
