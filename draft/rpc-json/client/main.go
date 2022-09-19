package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:11223")
	if err != nil {
		log.Fatal("dial failed: ", err)
	}
	defer conn.Close()

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	var reply string
	if err := client.Call("HelloService.Hi", "张三", &reply); err != nil {
		log.Fatal("rpc call failed: ", err)
	}

	log.Println("Result: ", reply)
}
