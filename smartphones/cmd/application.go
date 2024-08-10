package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	pb "github.com/slayersv/e-commerce/proto"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Application struct {
	ErrorLogger *log.Logger
	Infologger  *log.Logger
	HttpServer  *http.Server
	GrpcServer  *grpc.Server
	GrpcHandler *grpcHandler
	Kafka       *kgo.Client
	DB          *PostgresDB
}

func NewApplication() *Application {
	errorLogger := log.New(os.Stderr, "ERROR\t", log.Flags()|log.LUTC)
	infoLogger := log.New(os.Stdout, "INFO\t", log.Flags()|log.LUTC)

	conn, err := os.ReadFile("../DBConnectionString")
	if err != nil {
		errorLogger.Fatal(err)
	}
	DBConnString := string(conn)
	db, err := sql.Open("postgres", DBConnString)
	postgres := PostgresDB{db}
	if err != nil {
		errorLogger.Fatal(err)
	}
	httpserver := &http.Server{
		Addr:     "localhost:8080",
		ErrorLog: errorLogger,
	}

	grpcserver := grpc.NewServer()
	reflection.Register(grpcserver)
	grpchandler := &grpcHandler{
		Port:        8081,
		DB:          &postgres,
		ErrorLogger: errorLogger,
	}
	pb.RegisterSmartphoneServiceServer(grpcserver, grpchandler)
	app := &Application{
		ErrorLogger: errorLogger,
		Infologger:  infoLogger,
		HttpServer:  httpserver,
		GrpcServer:  grpcserver,
		GrpcHandler: grpchandler,
		DB:          &postgres,
	}
	app.HttpServer.Handler = app.NewRouter()

	kafka, err := app.NewKafkaClient()
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}
	app.Kafka = kafka
	app.Infologger.Println(app.Kafka.Ping(context.Background()))
	return app
}
