package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "realmadrid/protos"

	"google.golang.org/grpc"
)

const (
	port = ":8888"
)

var players []*pb.Player

type playerServer struct {
	pb.UnimplementedPlayerServiceServer
}

func main() {
	fmt.Println("grpc server running")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPlayerServiceServer(s, &playerServer{})
	log.Printf("server listening at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("err in serving: %v", err)
	}
}

func (s *playerServer) AddPlayer(ctx context.Context, req *pb.AddPlayerRequest) (*pb.AddPlayerResponse, error) {
	if req.Player.Id == 0 {
		return nil, fmt.Errorf("empty player ID")
	}
	plyr := &pb.Player{
		Id:          req.Player.Id,
		Name:        req.Player.Name,
		Nationality: req.Player.Nationality,
		Position:    req.Player.Position,
	}

	players = append(players, plyr)
	return &pb.AddPlayerResponse{
		Player: plyr,
	}, nil
}

func (s *playerServer) GetPlayer(ctx context.Context, req *pb.GetPlayerRequest) (*pb.GetPlayerResponse, error) {
	resp := &pb.Player{}
	for _, player := range players {
		if player.GetId() == req.GetId() {
			resp = player
		}
	}

	return &pb.GetPlayerResponse{
		Player: resp,
	}, nil
}
