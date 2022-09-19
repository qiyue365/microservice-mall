package main

import (
	"log"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:11223")
	if err != nil {
		log.Fatal("dial failed: ", err)
	}
	defer conn.Close()

	var reply string
	if err := conn.Call("HelloService.Hi", "张三", &reply); err != nil {
		log.Fatal("rpc call failed: ", err)
	}

	log.Println("Result: ", reply)
}
