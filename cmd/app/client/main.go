package main

import (
	"context"
	"fmt"
	"log"
	pb "realmadrid/protos"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cannot connect to server: %v", err)
	}
	defer conn.Close()
	client := pb.NewPlayerServiceClient(conn)
	addPlayer(client, &pb.Player{
		Id:          1,
		Name:        "Fran",
		Nationality: "Spain",
		Position:    "LB",
	})
	getPlayer(client, int64(1))
}

func getPlayer(client pb.PlayerServiceClient, id int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.GetPlayerRequest{Id: id}
	player, err := client.GetPlayer(ctx, req)
	if err != nil {
		log.Fatalf("err: %v in get player from client %v \n", err, client)
	}
	fmt.Println("player list:", player)
}

func addPlayer(client pb.PlayerServiceClient, plyr *pb.Player) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.AddPlayerRequest{Player: plyr}
	resp, err := client.AddPlayer(ctx, req)
	if err != nil {
		log.Fatalf("err: %v in add player from client %v \n", client, err)
	}
	if resp.GetPlayer() != nil {
		fmt.Println("player added:", resp)
	} else {
		fmt.Println("player not added")
	}
}
