package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/test", a.getProducts).Methods("GET")
	a.Router.HandleFunc("/todolist", a.gettodolistitems).Methods("GET")
	a.Router.HandleFunc("/todolist", a.addtodolist).Methods("POST")
}
