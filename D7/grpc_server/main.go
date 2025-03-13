package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	svc "hacktiv/handler/grpc"
	userPB "hacktiv/pb/user"
)

func main() {
	port := ":8080" // pake env
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	userSvc := svc.NewUserGrpcServer()
	gprcServer := grpc.NewServer()

	userPB.RegisterUserServiceServer(gprcServer, userSvc)

	fmt.Println("Server started at port ", port)
	if err := gprcServer.Serve(lis); err != nil {
		panic(err)
	}
}
