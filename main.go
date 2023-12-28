package main

import (
	"fmt"
	"github.com/Includeoyasi/cat/api"
	proto "github.com/Includeoyasi/cat/pkg/cat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port = "50051"

func main() {
	lst, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error init listener on port: %s, err: %s", port, err)
		return
	}

	server := grpc.NewServer()
	reflection.Register(server)
	proto.RegisterCatServer(server, api.GrpcServer{})

	log.Printf("server listening at: %s", lst.Addr())

	if err = server.Serve(lst); err != nil {
		log.Fatalf("Error serve; err: %s", err)
		return
	}
}
