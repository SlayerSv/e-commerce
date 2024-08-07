package main

import (
	"net/http"
	"runtime/debug"
)

func (app *Application) LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.Infologger.Println(r.URL, r.Method, r.RemoteAddr, r.Body)
		next.ServeHTTP(w, r)
	})
}

func (app *Application) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				app.ErrorJSON(w, errInternal)
				app.ErrorLogger.Println(err, string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, r)
	})
}
