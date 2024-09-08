package product

import (
	"context"
	"grpc-golang/pb"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	// Giả sử logic lấy product từ repository
	return &pb.GetProductResponse{
		ProductId: req.GetProductId(),
		Name:      "Sample Product",
		Price:     49.99,
	}, nil
}
