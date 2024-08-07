package main

import (
	"net/http"
)

func (app *Application) NewRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /smartphones", app.GetAll)
	router.HandleFunc("GET /smartphones/{id}", app.GetOne)
	router.HandleFunc("DELETE /smartphones/{id}", app.Delete)
	router.HandleFunc("POST /smartphones", app.Create)
	router.HandleFunc("UPDATE /smartphones", app.Update)
	return app.LogRequests(router)
}
