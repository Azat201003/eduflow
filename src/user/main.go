package main

import (
	"user-service/server"

	"user-service/config"

	pb "github.com/Azat201003/eduflow_service_api/gen/user"

	// server "eduflow/src/proto/user/server"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const SERVICE_ID = 0

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error with getting config: %v", err)
	}
	var service_config *config.Service = conf.Service
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", service_config.Host, service_config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, server.NewServer())
	go grpcServer.Serve(lis)
	fmt.Println("abeme")
}
