package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/qiyue365/microservice-mall/draft/grpc-interpretor/defs"
	"github.com/qiyue365/microservice-mall/draft/grpc-interpretor/pb"
	"google.golang.org/grpc"
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
		start := time.Now()
		defer func(start time.Time) {
			log.Println("服务端", info.FullMethod, "耗时：", time.Since(start))
		}(start)
		return handler(ctx, req)
	})

	srv := grpc.NewServer(opt)
	pb.RegisterGreeterServiceServer(srv, &GreeterServer{})
	log.Fatal(srv.Serve(lis))
}
