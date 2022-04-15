package util

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGRPC(addr string, block bool) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*2000)
	defer cancel()

	return dialGRPCContext(ctx, addr, block)
}

func DialGRPCContext(ctx context.Context, addr string, block bool) (*grpc.ClientConn, error) {
	return dialGRPCContext(ctx, addr, block)
}

func dialGRPCContext(ctx context.Context, addr string, block bool) (*grpc.ClientConn, error) {
	opts := make([]grpc.DialOption, 0, 2)
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if block {
		opts = append(opts, grpc.WithBlock())
	}

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		if conn != nil {
			_ = conn.Close()
		}
		return nil, err
	}

	return conn, nil
}
