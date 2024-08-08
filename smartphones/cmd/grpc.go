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

func (handler *grpcHandler) GetOne(ctx context.Context, request *pb.OneRequest) (*pb.OneResponse, error) {
	id := request.GetId()
	sm, err := handler.DB.GetOne(int(id))
	if err != nil {
		handler.ErrorLogger.Println(err)
		return nil, err
	}
	return &pb.OneResponse{
		Smartphone: handler.Convert(sm),
	}, nil
}

func (handler grpcHandler) GetMany(ctx context.Context, request *pb.ManyRequest) (*pb.ManyResponse, error) {
	smarts, err := handler.DB.GetAll()
	if err != nil {
		handler.ErrorLogger.Println(err)
		return nil, err
	}
	smartphones := make([]*pb.Smartphone, len(smarts))
	for i, sm := range smarts {
		smartphones[i] = handler.Convert(sm)
	}
	return &pb.ManyResponse{
		Smartphones: smartphones,
	}, nil
}

func (handler grpcHandler) Convert(sm Smartphone) *pb.Smartphone {
	return &pb.Smartphone{
		Id:          uint32(sm.ID),
		Model:       sm.Model,
		Producer:    sm.Producer,
		Color:       sm.Color,
		ScreenSize:  sm.ScreenSize,
		Description: sm.Description,
		Image:       sm.Image,
		Price:       uint32(sm.Price),
	}
}
