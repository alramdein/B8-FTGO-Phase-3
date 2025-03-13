package grpc

import (
	"context"

	userPB "github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/pb/user"
	"github.com/alramdein/B8-FTGO-Phase-3/D8-demo-grpc/user-service/usecase"
)

type server struct {
	userPB.UnimplementedUserServiceServer
	userUsecase usecase.IUserUsecase
}

func NewUserGrpcServer(userUsecase usecase.IUserUsecase) userPB.UserServiceServer {
	return &server{
		userUsecase: userUsecase,
	}
}

func (s *server) ListUsers(ctx context.Context, in *userPB.ListUsersRequest) (*userPB.ListUsersResponse, error) {
	users, err := s.userUsecase.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	var usersPB []*userPB.User
	for _, user := range users {
		usersPB = append(usersPB, &userPB.User{
			Id:    int32(user.ID),
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return &userPB.ListUsersResponse{
		Users: usersPB,
	}, nil
}

func (s *server) CreateUser(ctx context.Context, in *userPB.CreateUserRequest) (*userPB.User, error) {
	return &userPB.User{
		Id:    1,
		Name:  "John Doe",
		Email: "t0o0x@example.com",
	}, nil
}
