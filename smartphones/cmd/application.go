package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

type Application struct {
	ErrorLogger *log.Logger
	Infologger  *log.Logger
	Server      *http.Server
	DB          PostgresDB
}

func NewApplication() *Application {
	errorLogger := log.New(os.Stderr, "ERROR\t", log.LUTC|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO\t", log.LUTC)

	conn, err := os.ReadFile("DBConnectionString")
	if err != nil {
		errorLogger.Fatal(err)
	}
	DBConnString := string(conn)
	db, err := sql.Open("postgres", DBConnString)
	postgres := PostgresDB{db}
	if err != nil {
		errorLogger.Fatal(err)
	}
	server := &http.Server{
		Addr:     "localhost:8080",
		ErrorLog: errorLogger,
	}

	app := &Application{
		ErrorLogger: errorLogger,
		Infologger:  infoLogger,
		Server:      server,
		DB:          postgres,
	}

	app.Server.Handler = app.NewRouter()

	return app
}
