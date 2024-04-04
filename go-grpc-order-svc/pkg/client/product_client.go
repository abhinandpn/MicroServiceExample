package client

import (
	"context"
	"fmt"

	"github.com/abhinandpn/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ProductServiceClient struct {
	client pb.ProductServiceClient
}

func InitProductServiceClient(url string) ProductServiceClient {
	cc, err := grpc.Dial("tcp", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to dial")
	}
	c := ProductServiceClient{
		client: pb.NewProductServiceClient(cc),
	}
	return c
}

func (c *ProductServiceClient) FindOne(ProductId int) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: int64(ProductId),
	}
	return c.client.FindOne(context.Background(), req)
}

func (c *ProductServiceClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}
	return c.client.DecreaseStock(context.Background(), req)
}
