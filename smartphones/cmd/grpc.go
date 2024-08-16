package main

import (
	"context"
	"database/sql"
	"errors"
	"log"

	pb "github.com/slayersv/e-commerce/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			return nil, status.Error(codes.Internal, "internal error")
		}

	}
	return &pb.OneResponse{
		Smartphone: handler.Convert(sm),
	}, nil
}

func (handler grpcHandler) GetMany(ctx context.Context, request *pb.ManyRequest) (*pb.ManyResponse, error) {
	smarts, err := handler.DB.GetAll()
	if err != nil {
		handler.ErrorLogger.Println(err)
		return nil, status.Error(codes.Internal, "internal error")
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
