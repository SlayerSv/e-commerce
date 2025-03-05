package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var errInternal = errors.New("internal server error")
var errIncorrectCredentialsFormat = errors.New("incorrect credentials format")
var errInvalidCredentials = errors.New("invalid credentials")
var errUserExists = errors.New("user already exists")
var errUserNotFound = errors.New("user not found")
var errWrongPassword = errors.New("wrong password")

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.ErrorJSON(w, r, errIncorrectCredentialsFormat)
	}
	userDB, err := app.DB.GetUserByName(user.Name)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	if userDB.Password != user.Password {
		app.ErrorJSON(w, r, errWrongPassword)
		return
	}
	app.Encode(w, r, user)
}

func (app *Application) Signup(w http.ResponseWriter, r *http.Request) {
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		app.ErrorJSON(w, r, errIncorrectCredentialsFormat)
		return
	}
	err = app.DB.CreateUser(user)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (app *Application) ErrorJSON(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	app.ErrorLogger.Println(r.Method, r.URL, err.Error())
	var code int
	if errors.Is(err, errUserNotFound) || errors.Is(err, errWrongPassword) {
		err = errInvalidCredentials
		code = http.StatusUnauthorized
	} else if errors.Is(err, errIncorrectCredentialsFormat) {
		err = errIncorrectCredentialsFormat
		code = http.StatusBadRequest
	} else {
		err = errInternal
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}

func (app *Application) Encode(w http.ResponseWriter, r *http.Request, obj any) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(obj)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
}
