package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type todo struct {
	ID       int    `json:"id"`
	Todolist string `json:"todolist"`
}

func gettodolist(db *sql.DB, start, count int) ([]todo, error) {
	rows, err := db.Query(
		"SELECT id, todolist FROM todo ")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todolistitems := []todo{}

	for rows.Next() {
		var p todo
		if err := rows.Scan(&p.ID, &p.Todolist); err != nil {
			return nil, err
		}
		todolistitems = append(todolistitems, p)
	}

	return todolistitems, nil
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	http.ListenAndServe(":8010", a.Router)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	data_test := []string{}
	respondWithJSON(w, http.StatusOK, data_test)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func (a *App) gettodolistitems(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := gettodolist(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/test", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/todolist", a.gettodolistitems).Methods("GET")

}

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"faisal",
		"postgres")

	a.Run(":8010")
}
