package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	svc "hacktiv/handler/grpc"
	userPB "hacktiv/pb/user"
)

func main() {
	port := ":8080" // pake env
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(WithAuth),
	}

	userSvc := svc.NewUserGrpcServer()
	gprcServer := grpc.NewServer(opts...)

	userPB.RegisterUserServiceServer(gprcServer, userSvc)

	fmt.Println("Server started at port ", port)
	if err := gprcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func WithAuth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	if len(md["authorization"]) == 0 {
		return nil, fmt.Errorf("missing token")
	}

	// validate token blabla
	token, ok := md["authorization"]
	if !ok {
		return nil, fmt.Errorf("missing token")
	}

	if token[0] != "kajnfjanfandfkljadnALIF" {
		return nil, fmt.Errorf("invalid token")
	}

	return handler(ctx, req)
}
