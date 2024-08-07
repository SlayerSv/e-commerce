package main

import (
	"net"

	"google.golang.org/grpc"
)

func main() {
	app := NewApplication()
	go func() {

		app.Infologger.Println("starting server on address " + app.Server.Addr)
		app.ErrorLogger.Fatalln(app.Server.ListenAndServe())
	}()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}
	server := grpc.NewServer()
}
