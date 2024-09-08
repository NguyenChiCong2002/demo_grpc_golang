package order

import (
	"context"
	"grpc-golang/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	ProductClient pb.ProductServiceClient
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	// Gọi ProductService để lấy thông tin sản phẩm
	productReq := &pb.GetProductRequest{ProductId: req.GetProductId()}
	productRes, err := s.ProductClient.GetProduct(ctx, productReq)
	if err != nil {
		return nil, err
	}

	// Tính toán tổng tiền
	totalAmount := productRes.GetPrice() * float32(req.GetQuantity())

	return &pb.CreateOrderResponse{
		OrderId:     "order_1",
		Status:      "SUCCESS",
		TotalAmount: totalAmount,
	}, nil
}
