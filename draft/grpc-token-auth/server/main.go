package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/qiyue365/microservice-mall/draft/grpc-token-auth/defs"
	"github.com/qiyue365/microservice-mall/draft/grpc-token-auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {

	return &pb.HelloReply{
		Message: fmt.Sprintf("你好，%s", req.Name),
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", defs.GrpcServerAddr)
	if err != nil {
		log.Fatal("listen failed: ", err)
	}

	// 一元拦截器
	opt := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "no metadata")
		}
		tokenSlice, ok := md[defs.AuthTokenName]
		if !ok {
			return nil, status.Error(codes.InvalidArgument, "no token")
		}
		if len(tokenSlice) != 1 {
			return nil, status.Error(codes.InvalidArgument, "invalid token slice")
		}
		token := tokenSlice[0]
		if token != "yUUcePTbfC3h84xqkV27nLGQ" {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}
		return handler(ctx, req)
	})

	srv := grpc.NewServer(opt)
	pb.RegisterGreeterServiceServer(srv, &GreeterServer{})
	log.Fatal(srv.Serve(lis))
}
