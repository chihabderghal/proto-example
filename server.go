package main

import (
	"context"
	"log"
	"net"

	pb "github.com/chihabderghal/proto-example/coffeeshop_proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *server) GetMenu(menuRequest *pb.MenuRequest, srv pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item {
		&pb.Item{
			Id: "1",
			Name: "Black Coffe",
		},
		&pb.Item{
			Id: "2",
			Name: "Americano",
		},
		&pb.Item{
			Id: "3",
			Name: "Vanilla Soy Chai Latte",
		},
	}

	for i, _ := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}
	
	return nil
}

func (s *server) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "ABC123",
	}, nil
}

func (s *server) GetOrderStatus(context context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderIid: receipt.Id,
		Status: "In Progress",
	}, nil
}

func main() {
	list, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("filed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &server{})

	if err := grpcServer.Serve(list); err != nil {
		log.Fatal("filed to serve: %s", err)

	}
}
