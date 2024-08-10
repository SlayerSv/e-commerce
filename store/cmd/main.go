package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/slayersv/e-commerce/proto"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	connection, err := grpc.NewClient("localhost:8081", opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer connection.Close()
	client := pb.NewSmartphoneServiceClient(connection)
	ctx := context.Background()
	smartphone, err := client.GetOne(ctx, &pb.OneRequest{Id: 1})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(smartphone)
}
