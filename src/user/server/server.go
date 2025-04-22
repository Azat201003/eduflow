package server

import (
	"context"

	pb "github.com/Azat201003/eduflow_service_api/gen/user"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (*userServiceServer) Login(context context.Context, creditionals *pb.Creditionals) (*pb.Token, error) {
	return &pb.Token{Token: "abeme"}, nil
}

func (*userServiceServer) Register(context.Context, *pb.Creditionals) (*pb.Token, error) {
	return &pb.Token{Token: "abeme"}, nil
}

func (*userServiceServer) GetById(ctx context.Context, id *pb.Id) (*pb.User, error) {
	return &pb.User{Id: 0, Username: "John"}, nil
}

func NewServer() pb.UserServiceServer {
	return new(userServiceServer)
}
