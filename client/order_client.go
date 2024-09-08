package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-golang/pb"
	"log"
)

func main() {
	// Kết nối tới OrderService
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to OrderService: %v", err)
	}
	defer conn.Close()

	orderClient := pb.NewOrderServiceClient(conn)

	// Gửi yêu cầu tạo đơn hàng
	req := &pb.CreateOrderRequest{
		ProductId: "product_1",
		Quantity:  2,
	}
	res, err := orderClient.CreateOrder(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while creating order: %v", err)
	}

	log.Printf("Order created: %v", res)
}
