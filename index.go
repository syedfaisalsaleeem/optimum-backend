package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS todo
(
    id SERIAL,
    Todolist TEXT
)`

func (a *App) Run(addr string) {
	http.ListenAndServe(addr, a.Router)
}

func (a *App) Initialize(user, password, host, dbname string) {
	// setting up the postgres database and initializing the endpoints

	connectionString := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", user, password, host, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.DB.Exec(tableCreationQuery)
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func main() {
	// Main Function loads when app is started
	a := App{}
	// get env from loadenv.go
	a.Initialize(
		goDotEnvVariable("DB_USER"),
		goDotEnvVariable("DB_PASSWORD"),
		goDotEnvVariable("DB_HOST"),
		goDotEnvVariable("DB_NAME"),
	)
	// start the http server
	a.Run(":8010")
}
