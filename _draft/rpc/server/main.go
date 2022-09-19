package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hi(req string, reply *string) error {
	*reply = fmt.Sprintf("你好，%s", req)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":11223")
	if err != nil {
		log.Fatal("listen failed: ", err)
	}
	if err := rpc.RegisterName("HelloService", &HelloService{}); err != nil {
		log.Fatal("rpc register failed: ", err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal("accept failed: ", err)
		}
		go rpc.ServeConn(conn)
	}
}
