package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hi(req string, reply *string) error {
	*reply = fmt.Sprintf("你好，%s", req)
	return nil
}

func main() {

	if err := rpc.RegisterName("HelloService", &HelloService{}); err != nil {
		log.Fatal("rpc register failed: ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.ReadCloser
			io.Writer
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":11223", nil)
}
