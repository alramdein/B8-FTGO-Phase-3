package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func NewGRPCClient(address string) (*grpc.ClientConn, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(withAuth),
	}

	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	return conn, nil
}

func withAuth(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "123")
	return invoker(newCtx, method, req, reply, cc, opts...)
}
