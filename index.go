package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {
	http.ListenAndServe(":8010", a.Router) 
}

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"postgres",
		"testing")
	a.Run(":8010")
}