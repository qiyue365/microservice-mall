package main

import (
	"context"
	"log"

	"github.com/qiyue365/microservice-mall/draft/grpc-token-auth/defs"
	"github.com/qiyue365/microservice-mall/draft/grpc-token-auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TokenCredential struct{}

func (c *TokenCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		defs.AuthTokenName: defs.AuthToken,
	}, nil
}
func (c *TokenCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// 客户端一元拦截器
	// opt := grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	md := metadata.Pairs(defs.AuthTokenName, defs.AuthToken)
	// 	ctx = metadata.NewOutgoingContext(ctx, md)
	// 	return invoker(ctx, method, req, reply, cc, opts...)
	// })
	opt := grpc.WithPerRPCCredentials(&TokenCredential{})
	conn, err := grpc.Dial(defs.GrpcServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), opt)
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
