package main

import (
	"fmt"
	"net"
)

func main() {
	app := NewApplication()
	go func() {

		app.Infologger.Println("starting server on address " + app.HttpServer.Addr)
		app.ErrorLogger.Fatalln(app.HttpServer.ListenAndServe())
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.GrpcHandler.Port))
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}
	app.Infologger.Println("listening grpc requests at ", lis.Addr())
	app.ErrorLogger.Fatalln(app.GrpcServer.Serve(lis))

}
