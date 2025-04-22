package main

import (
	"user-service/server"

	config "github.com/Azat201003/eduflow_service_api/config"

	pb "github.com/Azat201003/eduflow_service_api/gen/user"

	// server "eduflow/src/proto/user/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const SERVICE_ID = 0

func main() {
	// config getting
	conf, err := config.GetConfig("../../config.yaml")
	if err != nil {
		log.Fatalf("Error with getting config: %v", err)
	}
	service_config, err := conf.GetServiceById(0)
	if err != nil {
		log.Fatalf("Error with finding condfig: %v", err)
	}

	// database connecting
	db_conf := conf.Database
	dsn := fmt.Sprintf("host=%v user=%v password=1234 dbname=%v port=%v sslmode=disable TimeZone=Europe/Moscow", db_conf.Host, service_config.DB_user, service_config.DB, db_conf.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with connecting to db: %v", err)
	}

	// starting server
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", service_config.Host, service_config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, server.NewServer(db))

	fmt.Println(lis.Addr())
	grpcServer.Serve(lis)
}
