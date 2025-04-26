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

func (s *userServiceServer) Login(context context.Context, creditionals *pb.Creditionals) (*pb.Token, error) {
	user := db.User{Username: creditionals.Username, Password: creditionals.Password}
	err := db.FindUser(s.db, &user)
	return &pb.Token{Token: user.Token}, err
}

func (s *userServiceServer) Register(context context.Context, creditionals *pb.Creditionals) (*pb.Token, error) {
	user := db.User{Username: creditionals.Username, Password: creditionals.Password}
	err := db.CreateUser(s.db, &user)
	return &pb.Token{Token: user.Token}, err
}

func (s *userServiceServer) GetUserById(ctx context.Context, id *pb.Id) (*pb.User, error) {
	user := db.User{ID: uint64(id.Id)}
	db.FindUser(s.db, &user)
	return &pb.User{Id: id, Username: user.Username}, nil
}

func (s *userServiceServer) GetUserByToken(ctx context.Context, token *pb.Token) (*pb.User, error) {
	user := db.User{Token: token.Token}
	db.FindUser(s.db, &user)
	id := pb.Id{Id: user.ID}
	return &pb.User{Id: &id, Username: user.Username}, nil
}

func NewServer(db *gorm.DB) pb.UserServiceServer {
	server := new(userServiceServer)
	server.db = db
	return server
}
