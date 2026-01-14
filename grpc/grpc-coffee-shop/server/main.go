package main

import (
	"context"
	"log"
	"net"

	pb "github.com/shreyansh28801/grpc-coffee-shop/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(
	req *pb.MenuRequest,
	stream pb.CoffeeShop_GetMenuServer,
) error {

	items := []*pb.Item{
		{Id: "1", Name: "Black Coffee"},
		{Id: "2", Name: "Americano"},
		{Id: "3", Name: "Vanilla Latte"},
	}

	for i := range items {
		stream.Send(&pb.Menu{
			Items: items[:i+1],
		})
	}
	return nil
}

func (s *server) PlaceOrder(
	ctx context.Context,
	order *pb.Order,
) (*pb.Receipt, error) {

	return &pb.Receipt{Id: "ORDER123"}, nil
}

func (s *server) GetOrderStatus(
	ctx context.Context,
	receipt *pb.Receipt,
) (*pb.OrderStatus, error) {

	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN_PROGRESS",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	log.Println("Server running on :9001")
	grpcServer.Serve(lis)
}
