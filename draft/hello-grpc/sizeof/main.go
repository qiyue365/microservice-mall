package main

import (
	"encoding/json"
	"log"

	"github.com/qiyue365/microservice-mall/draft/hello-grpc/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	req := &pb.HelloRequest{
		Name: "张三",
	}
	buf, err := proto.Marshal(req)
	if err != nil {
		log.Fatal("proto marshal failed: ", err)
	}
	log.Printf("%X, %d\n", buf, len(buf))

	buf, err = json.Marshal(req)
	if err != nil {
		log.Fatal("json marshal failed: ", err)
	}
	log.Printf("%X, %d\n", buf, len(buf))
}
