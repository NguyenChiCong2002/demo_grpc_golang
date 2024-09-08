package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc-golang/internal/product"
	"grpc-golang/pb"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &product.ProductService{})

	log.Println("ProductService is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
