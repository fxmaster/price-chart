package grpc

import (
	"context"
	"log"
	"net"
	pb "price-chart/pkg/protobuf/client"
	"price-chart/pkg/util"

	"google.golang.org/grpc"
)

func Init(ctx context.Context, addr string, mongo util.Mongo) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("failed to listen address: " + err.Error())
	}

	log.Println("Start server...")
	server := grpc.NewServer()

	pb.RegisterClientServiceServer(server, &Server{})

	ch := make(chan error)

	go func() {
		defer close(ch)

		ch <- server.Serve(listener)
	}()

	select {
	case <-ch:
		log.Println("failed to serve address:" + addr)
	case <-ctx.Done():
		log.Println("server terminated")
		log.Println("waiting for all goroutines stop...")
		server.GracefulStop()
	}
}
