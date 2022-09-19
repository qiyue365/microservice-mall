package main

import (
	"context"
	"log"
	"time"

	"github.com/qiyue365/microservice-mall/draft/grpc-metadata/defs"
	"github.com/qiyue365/microservice-mall/draft/grpc-metadata/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	md := metadata.Pairs("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	conn, err := grpc.Dial(defs.GrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("grpc dail failed: ", err)
	}
	defer conn.Close()

	client := pb.NewGreeterServiceClient(conn)
	reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "张三"})
	if err != nil {
		log.Fatal("grpc call failed: ", err)
	}
	log.Println(reply.Message)
}
