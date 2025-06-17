package main

import (
	"summary-service/server"
	"summary-service/server/validators"

	config "github.com/Azat201003/eduflow/backend/libs/config"

	"fmt"
	"log"
	"net"

	pb "github.com/Azat201003/eduflow/backend/libs/gen/go/summary"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const SERVICE_ID = 1

func main() {
	// config getting
	conf, err := config.GetConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Error with getting config: %v", err)
	}
	service_config, err := conf.GetServiceById(SERVICE_ID)
	if err != nil {
		log.Fatalf("Error with finding condfig: %v", err)
	}

	// database connecting
	db_conf := conf.Database
	conn_conf := service_config.Connect
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Moscow search_path=%v", db_conf.Host, conn_conf.User, conn_conf.Password, conn_conf.DB, db_conf.Port, conn_conf.Schema)
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
	pb.RegisterSummaryServiceServer(grpcServer, server.NewServer(db, new(validators.Validator)))

	fmt.Println(lis.Addr())
	grpcServer.Serve(lis)
}
