package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	"github.com/qiyue365/microservice-mall/draft/grpc-metadata/defs"
	"github.com/qiyue365/microservice-mall/draft/grpc-metadata/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata")
	}
	for key, value := range md {
		log.Println(key, value)
	}
	ts, ok := md["timestamp"]
	if !ok {
		return nil, errors.New("no timestamp")
	}
	return &pb.HelloReply{
		Message: fmt.Sprintf("你好，%s @ %q", req.Name, ts),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", defs.GrpcServerAddr)
	if err != nil {
		log.Fatal("listen failed: ", err)
	}
	srv := grpc.NewServer()
	pb.RegisterGreeterServiceServer(srv, &GreeterServer{})
	log.Fatal(srv.Serve(lis))
}
