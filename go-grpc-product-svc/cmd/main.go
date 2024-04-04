package main

import (
	"fmt"
	"log"
	"net"

	"github.com/abhinandpn/go-grpc-product-svc/pkg/config"
	"github.com/abhinandpn/go-grpc-product-svc/pkg/db"
	"github.com/abhinandpn/go-grpc-product-svc/pkg/pb"
	"github.com/abhinandpn/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("failed to config", err)
	}
	h := db.Init(c.DBurl)
	list, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	fmt.Println("product Svc on", c.Port)
	s := services.Server{
		H: h,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalln("failed to serve", err)
	}
}
