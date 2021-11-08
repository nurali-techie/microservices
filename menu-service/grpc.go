package main

import (
	"context"
	"net"

	pb "github.com/nurali-techie/microservices/commons-go/proto"
	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

type MenuServiceServer struct {
	pb.UnimplementedMenuServiceServer

	repo *RestoRepo
}

func NewMenuServiceServer(repo *RestoRepo) *MenuServiceServer {
	return &MenuServiceServer{
		repo: repo,
	}
}

func (s *MenuServiceServer) Start() {
	l, err := net.Listen("tcp", ":2012")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterMenuServiceServer(server, s)
	go func() {
		log.Fatal(server.Serve(l))
	}()
}

func (s *MenuServiceServer) GetRestaurant(ctx context.Context, req *pb.GetRestaurantRequest) (*pb.GetRestaurantReply, error) {
	resto, err := s.repo.GetResto(req.ID)
	if err != nil {
		return nil, err
	}
	return &pb.GetRestaurantReply{
		ID:      resto.ID,
		Name:    resto.Name,
		Address: resto.Address,
		City:    resto.City,
	}, nil
}
