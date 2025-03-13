package grpc

import (
	"context"
	userPB "hacktiv/pb/user"
)

type server struct {
	userPB.UnimplementedUserServiceServer
}

func NewUserGrpcServer() userPB.UserServiceServer {
	return &server{}
}

func (s *server) ListUsers(ctx context.Context, in *userPB.ListUsersRequest) (*userPB.ListUsersResponse, error) {
	return &userPB.ListUsersResponse{
		Users: []*userPB.User{
			{
				Id:    1,
				Name:  "Alif 123",
				Email: "t0o0x@example.com",
			},
			{
				Id:    2,
				Name:  "Jane Faq",
				Email: "t0o0x@example.com",
			},
		},
	}, nil
}

func (s *server) CreateUser(ctx context.Context, in *userPB.CreateUserRequest) (*userPB.User, error) {
	return &userPB.User{
		Id:    1,
		Name:  "John Doe",
		Email: "t0o0x@example.com",
	}, nil
}
