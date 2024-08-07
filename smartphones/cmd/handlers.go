package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var errNotfound = errors.New("not found")
var errInternal = errors.New("internal server error")
var errIncorrectID = errors.New("incorrect id")

func (app *Application) GetAll(w http.ResponseWriter, r *http.Request) {
	smartphones, err := app.DB.GetAll()
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	app.Encode(w, smartphones)
}

func (app *Application) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := app.ExtractID(r)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	smartphone, err := app.DB.GetOne(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	app.Encode(w, smartphone)
}

func (app *Application) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := app.ExtractID(r)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	smartphone, err := app.DB.Delete(id)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	app.Encode(w, smartphone)
}

func (app *Application) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	sm := Smartphone{}
	decoder.Decode(&sm)
	sm, err := app.DB.Update(sm)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	app.Encode(w, sm)
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	sm := Smartphone{}
	decoder.Decode(&sm)
	sm, err := app.DB.Create(sm)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
	app.Encode(w, sm)
}

func (app *Application) ErrorJSON(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	app.ErrorLogger.Println(err)
	var code int
	switch err {
	case sql.ErrNoRows:
		err = errNotfound
		code = http.StatusNotFound
	case errIncorrectID:
		code = http.StatusBadRequest
	default:
		err = errInternal
		code = http.StatusInternalServerError
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func (app *Application) Encode(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(obj)
	if err != nil {
		app.ErrorJSON(w, err)
		return
	}
}

func (app *Application) ExtractID(r *http.Request) (int, error) {
	stringId := r.PathValue("id")
	id, err := strconv.Atoi(stringId)
	if stringId == "" || err != nil {
		return 0, errIncorrectID
	}
	return id, nil
}
