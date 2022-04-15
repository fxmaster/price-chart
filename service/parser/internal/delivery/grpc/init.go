package grpc

import (
	"context"
	"log"
	"net"
)

func Init(ctx context.Context, addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("failed to listen address: " + err.Error())
	}

	log.Println("Start server...")
	log.Println(listener)
	// server := grpc.NewServer()

	// pb.RegisterUserReaderServiceServer(server, &Server{
	// 	Service: service.UserService{
	// 		Repo: model.UserRepository{
	// 			Mongo: mongo,
	// 		},
	// 	},
	// })

	// ch := make(chan error)

	// go func() {
	// 	defer close(ch)

	// 	ch <- server.Serve(listener)
	// }()

	// select {
	// case <-ch:
	// 	log.Println("failed to serve address:" + addr)
	// case <-ctx.Done():
	// 	log.Println("server terminated")
	// 	log.Println("waiting for all goroutines stop...")
	// 	server.GracefulStop()
	// }
}
