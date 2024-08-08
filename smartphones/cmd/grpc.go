package main

import (
	"context"
	"log"

	pb "github.com/slayersv/e-commerce/proto"
)

type grpcHandler struct {
	pb.UnimplementedSmartphoneServiceServer
	Port        int
	DB          *PostgresDB
	ErrorLogger *log.Logger
}

func (server *grpcHandler) GetOne(ctx context.Context, request *pb.OneRequest) (*pb.OneResponse, error) {
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
