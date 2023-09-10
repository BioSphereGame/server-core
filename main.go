package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/BioSphereGame/server-core/counter"
	pb "github.com/BioSphereGame/server-core/proto/counter"
)

var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCounterServiceServer(s, &counter.Counter{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
