package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"user-service/server/db"

	"github.com/Azat201003/eduflow/backend/libs/config"
	"github.com/Azat201003/eduflow/backend/libs/gen/go/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserTestSuite struct {
	suite.Suite
	Client *user.UserServiceClient
	dbm    *db.DBManger
}

func TestUserSuite(t *testing.T) {
	t.Helper()
	t.Parallel()
	// include configuration
	conf, err := config.GetConfig("../../../config.yaml")
	assert.NoError(t, err)

	user_conf, err := conf.GetServiceById(0)
	assert.NoError(t, err)

	// connecting to user service
	user_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", user_conf.Host, user_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)

	user_client := user.NewUserServiceClient(user_conn)

	// connecting db
	db_conf := conf.Database
	conn_conf := user_conf.Connect
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Moscow search_path=%v", db_conf.Host, conn_conf.User, conn_conf.Password, conn_conf.DB, db_conf.Port, conn_conf.Schema)
	db_conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with connecting to db: %v", err)
	}

	s := UserTestSuite{
		Client: &user_client,
		dbm:    &db.DBManger{DB: db_conn},
	}
	suite.Run(t, &s)
}

func (s *UserTestSuite) TestLoggingIn() {
	var opts []grpc.CallOption
	token, err := (*s.Client).Login(context.Background(), &user.Creditionals{Username: "Coolman", Password: "1234"}, opts...)
	s.NoError(err)
	s.Equal(token.Token, "anqrfsNqOu")
}

func (s *UserTestSuite) TestGettingByToken() {
	user_obj, err := (*s.Client).GetUserByToken(context.Background(), &user.Token{Token: "anqrfsNqOu"})
	s.NoError(err)
	s.Equal(user_obj.Username, "Coolman")
}

func (s *UserTestSuite) TestGettingById() {
	user_obj, err := (*s.Client).GetUserById(context.Background(), &user.Id{Id: 3})
	s.NoError(err)
	s.Equal(user_obj.Username, "Coolman")
}

func (s *UserTestSuite) TestByNonexistenTokenGetting() {
	user_obj, err := (*s.Client).GetUserByToken(context.Background(), &user.Token{Token: "Abeme don't rules the world"})
	s.Error(err)
	s.Nil(user_obj)
}

func (s *UserTestSuite) TestByNonexistenIdGetting() {
	user_obj, err := (*s.Client).GetUserById(context.Background(), &user.Id{Id: 525252})
	s.Error(err)
	s.Nil(user_obj)
}
