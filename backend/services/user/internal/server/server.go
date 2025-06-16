package server

import (
	"context"

	"user-service/server/db"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/user"
	"gorm.io/gorm"
)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
	dbm *db.DBManger
}

func (s *userServiceServer) Login(context context.Context, creditionals *pb.Creditionals) (*pb.Token, error) {
	user := db.User{Username: creditionals.Username, Password: creditionals.Password}
	err := s.dbm.FindUser(&user)
	return &pb.Token{Token: user.Token}, err
}

func (s *userServiceServer) Register(context context.Context, creditionals *pb.Creditionals) (*pb.Token, error) {
	user := db.User{Username: creditionals.Username, Password: creditionals.Password}
	err := s.dbm.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	err = s.dbm.FindUser(&user)
	return &pb.Token{Token: user.Token}, err
}

func (s *userServiceServer) GetUserById(ctx context.Context, id *pb.Id) (*pb.User, error) {
	user := db.User{ID: id.Id}
	err := s.dbm.FindUser(&user)
	return &pb.User{
		Id:       id,
		Username: user.Username,
		IsStaff:  user.IsStaff,
	}, err
}

func (s *userServiceServer) GetUserByToken(ctx context.Context, token *pb.Token) (*pb.User, error) {
	user := db.User{Token: token.Token}
	err := s.dbm.FindUser(&user)
	id := pb.Id{Id: user.ID}
	return &pb.User{
		Id:       &id,
		Username: user.Username,
		IsStaff:  user.IsStaff,
	}, err
}

func NewServer(db_conn *gorm.DB) pb.UserServiceServer {
	server := new(userServiceServer)
	server.dbm = &db.DBManger{DB: db_conn}
	return server
}
