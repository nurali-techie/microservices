package main

import (
	"context"

	pb "github.com/nurali-techie/microservices/commons-go/proto"
	"google.golang.org/grpc"
)

type MenuService struct {
	client pb.MenuServiceClient
}

func NewMenuClient() *grpc.ClientConn {
	conn, err := grpc.Dial("menu_service:2012", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return conn
}

func NewMenuService(client *grpc.ClientConn) *MenuService {
	return &MenuService{
		client: pb.NewMenuServiceClient(client),
	}
}

func (s *MenuService) GetRestaurant(restoID string) (*Restaurant, error) {
	req := pb.GetRestaurantRequest{ID: restoID}
	reply, err := s.client.GetRestaurant(context.Background(), &req)
	if err != nil {
		return nil, err
	}

	return &Restaurant{
		ID:      reply.ID,
		Name:    reply.Name,
		Address: reply.Address,
		City:    reply.City,
	}, nil
}
