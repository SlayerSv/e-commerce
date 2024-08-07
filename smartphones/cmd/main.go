package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/slayersv/e-commerce/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	pb.UnimplementedSmartphoneServiceServer
	DB          *PostgresDB
	ErrorLogger *log.Logger
}

func (server *grpcServer) GetOneGrpc(ctx context.Context, request *pb.OneRequest) (*pb.OneResponse, error) {
	id := request.GetId()
	sm, err := server.DB.GetOne(int(id))
	if err != nil {
		server.ErrorLogger.Println(err)
		return nil, err
	}
	return &pb.OneResponse{
		Smartphone: &pb.Smartphone{
			Id:          uint32(sm.ID),
			Model:       sm.Model,
			Producer:    sm.Producer,
			Color:       sm.Color,
			ScreenSize:  sm.ScreenSize,
			Description: *sm.Description,
			Image:       *sm.Image,
			Price:       uint32(sm.Price),
		},
	}, nil
}

const grpcPort = 8081

func main() {
	app := NewApplication()
	go func() {

		app.Infologger.Println("starting server on address " + app.Server.Addr)
		app.ErrorLogger.Fatalln(app.Server.ListenAndServe())
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}
	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterSmartphoneServiceServer(server, &grpcServer{
		DB:          app.DB,
		ErrorLogger: app.ErrorLogger,
	})
	app.Infologger.Println("listening grpc requests at ", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}
}
