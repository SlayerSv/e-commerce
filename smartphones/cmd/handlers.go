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
		app.ErrorJSON(w, r, err)
		return
	}
	app.Encode(w, r, smartphones)
}

func (app *Application) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := app.ExtractID(r)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	smartphone, err := app.DB.GetOne(id)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	app.Encode(w, r, smartphone)
}

func (app *Application) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := app.ExtractID(r)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	smartphone, err := app.DB.Delete(id)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	app.Encode(w, r, smartphone)
}

func (app *Application) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	sm := Smartphone{}
	decoder.Decode(&sm)
	sm, err := app.DB.Update(sm)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	app.Encode(w, r, sm)
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	sm := Smartphone{}
	decoder.Decode(&sm)
	sm, err := app.DB.Create(sm)
	if err != nil {
		app.ErrorJSON(w, r, err)
		return
	}
	app.Encode(w, r, sm)
}

func (app *Application) ErrorJSON(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	app.ErrorLogger.Println(r.Method, r.URL, err.Error())
	app.NewErrorMessage(err)
	var code int
	if errors.Is(err, sql.ErrNoRows) {
		err = errNotfound
		code = http.StatusNotFound
	} else if errors.Is(err, errIncorrectID) {
		err = errIncorrectID
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

func (app *Application) Encode(w http.ResponseWriter, r *http.Request, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(obj)
	if err != nil {
		app.ErrorJSON(w, r, err)
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
