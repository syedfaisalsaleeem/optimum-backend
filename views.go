package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	// test the endpoint is working or not

	data_test := []string{}
	respondWithJSON(w, http.StatusOK, data_test)
}

func (a *App) gettodolistitems(w http.ResponseWriter, r *http.Request) {
	// get the todoitems from the postgresql

	// setting up cors to enable all the users can access the endpoint
	setupCorsResponse(&w)

	products, err := gettodolist(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) addtodolist(w http.ResponseWriter, r *http.Request) {
	// adding the todoitems in the postgres database

	// setting up cors to enable all the users can access the endpoint
	setupCorsResponse(&w)

	var p todo
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	space := regexp.MustCompile(`\s+`)
	data := space.ReplaceAllString(p.Todolist, " ")
	if data == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if data == " " {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := p.createtodolistitem(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, p)
}
