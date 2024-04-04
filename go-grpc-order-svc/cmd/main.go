package main

import (
	"fmt"
	"log"
	"net"

	"github.com/abhinandpn/go-grpc-order-svc/pkg/client"
	"github.com/abhinandpn/go-grpc-order-svc/pkg/config"
	"github.com/abhinandpn/go-grpc-order-svc/pkg/db"
	"github.com/abhinandpn/go-grpc-order-svc/pkg/pb"
	"github.com/abhinandpn/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed at config %v", err)
	}
	h := db.Init(c.DB_Url)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}
	productSvc := client.InitProductServiceClient(c.ProductSvcUrl) //need to change
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}

}
