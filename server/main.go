package main

import (
	"fmt"
	protobuf "github.com/shukubota/protobuf-playground/grpc/v1"
	"github.com/shukubota/protobuf-playground/handlers"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("-------------aaa")
	gRPCServer := grpc.NewServer()

	// serviceを登録
	protobuf.RegisterHealthCheckServiceServer(gRPCServer, handlers.NewHealthCheckHandler())

	address := "127.0.0.1:7001"
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("err in listening")
	}

	m := cmux.New(listener)
	fmt.Printf("listening on %v\n", address)
	m.Serve()
}
