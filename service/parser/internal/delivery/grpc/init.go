package grpc

import (
	"context"
	"log"
	"net"
	pb "price-chart/pkg/protobuf/parser"
	"price-chart/service/parser/internal/service/handler"
	"price-chart/service/parser/internal/service/parser"

	"google.golang.org/grpc"
)

func Init(ctx context.Context, addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("failed to listen address: " + err.Error())
	}

	log.Println("Start server...")
	server := grpc.NewServer()

	pb.RegisterParserServiceServer(server, &Server{
		Handler: handler.GRPC{
			Parser: parser.Chromedp{},
		},
	})

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
