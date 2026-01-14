package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/shreyansh28801/grpc-coffee-shop/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:9001",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Stream menu
	stream, err := client.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatal(err)
	}

	var items []*pb.Item

	for {
		menu, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		items = menu.Items
		log.Println("Menu update:", menu.Items)
	}

	// Place order
	receipt, _ := client.PlaceOrder(ctx, &pb.Order{Items: items})
	log.Println("Receipt:", receipt.Id)

	// Get order status
	status, _ := client.GetOrderStatus(ctx, receipt)
	log.Println("Order status:", status.Status)
}
