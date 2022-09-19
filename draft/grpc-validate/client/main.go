package main

import (
	"context"
	"log"

	"github.com/qiyue365/microservice-mall/draft/grpc-validate/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:11223", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial failed:", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	rep, err := client.SayHello(ctx, &pb.Person{
		Id:     12345,
		Email:  "foo@bar.com",
		Mobile: "13800138000",
	})

	if err != nil {
		log.Fatal("call failed:", err)
	}

	log.Println(rep)
}
