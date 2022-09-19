package main

import (
	"context"
	"log"
	"time"

	"github.com/qiyue365/microservice-mall/draft/grpc-interpretor/defs"
	"github.com/qiyue365/microservice-mall/draft/grpc-interpretor/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	// 客户端一元拦截器
	opt := grpc.WithUnaryInterceptor(func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		defer func(start time.Time) {
			log.Println("客户端", method, "耗时：", time.Since(start))
		}(start)
		return invoker(ctx, method, req, reply, cc, opts...)
	})
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
