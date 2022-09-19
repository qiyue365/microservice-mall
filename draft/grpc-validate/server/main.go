package main

import (
	"context"
	"log"
	"net"

	"github.com/qiyue365/microservice-mall/draft/grpc-validate/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.Person) (*pb.Person, error) {
	return &pb.Person{
		Id:     req.Id,
		Email:  req.Email,
		Mobile: req.Mobile,
	}, nil
}

// Validator 验证器
type Validator interface {
	// Validate 验证
	Validate() error
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:11223")
	if err != nil {
		log.Fatal("listen failed:", err)
	}
	opt := grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r, ok := req.(Validator); ok {
			if err := r.Validate(); err != nil {
				return nil, status.Error(codes.InvalidArgument, err.Error())
			}
		}
		return handler(ctx, req)
	})
	srv := grpc.NewServer(opt)
	pb.RegisterGreeterServer(srv, &GreeterServer{})
	log.Fatal(srv.Serve(lis))
}
