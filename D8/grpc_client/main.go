package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "hacktiv/pb/user"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(WithAuth),
	}

	conn, err := grpc.NewClient("localhost:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	userServiceClient := pb.NewUserServiceClient(conn)

	// panggil fungsi service lain, seolah itu adalah fungsi dari service tsb
	_, err = userServiceClient.CreateUser(context.Background(), &pb.CreateUserRequest{
		Name:  "John Doe",
		Email: "t0o0x@example.com",
	})

	// fmt.Println("Create users")
	// fmt.Println(user, err)

	users, err := userServiceClient.ListUsers(context.Background(), &pb.ListUsersRequest{})

	fmt.Println("List users")
	fmt.Println(users, err)
}

func WithAuth(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	newCtx := metadata.AppendToOutgoingContext(ctx, "authorization", "123")
	return invoker(newCtx, method, req, reply, cc, opts...)
}
