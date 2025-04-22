package server

import (
	"context"

	"user-service/server/db"

	pb "github.com/Azat201003/eduflow_service_api/gen/user"
	"gorm.io/gorm"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	db *gorm.DB
}

func (*userServiceServer) Login(context context.Context, creditionals *pb.Creditionals) (*pb.Token, error) {
	return &pb.Token{Token: "abeme"}, nil
}

func (s *userServiceServer) Register(context.Context, *pb.Creditionals) (*pb.Token, error) {
	db.CreateUser(s.db)
	return &pb.Token{Token: "abeme"}, nil
}

func (*userServiceServer) GetById(ctx context.Context, id *pb.Id) (*pb.User, error) {
	return &pb.User{Id: 0, Username: "John"}, nil
}

func NewServer(db *gorm.DB) pb.UserServiceServer {
	server := new(userServiceServer)
	server.db = db
	return server
}
