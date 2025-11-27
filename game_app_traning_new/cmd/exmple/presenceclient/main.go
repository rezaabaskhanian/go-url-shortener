package main

import (
	"context"
	"fmt"
	presenceClient "game_app/internal/adapter/presence"
	"game_app/internal/contract/golang/presence"
	"game_app/internal/param"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial(":8086", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := presenceClient.New(conn)

	resp, err := client.GetPresence(context.Background(), param.GetPresenceRequest{UserIDs: []uint{1, 2, 4}})

	if err != nil {
		panic(err)
	}

	for _, item := range resp.Items {
		fmt.Println(item.UserID, item.TimeStamp)
	}
}
