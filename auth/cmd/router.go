package main

import (
	"net/http"
)

func (app *Application) NewRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("GET /login", app.Login)
	router.HandleFunc("POST /signup", app.Signup)
	return app.LogRequests(router)
}
