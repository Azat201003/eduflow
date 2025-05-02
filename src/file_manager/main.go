package main

import (
	"filager-service/server"

	config "github.com/Azat201003/eduflow_service_api/config"
	"github.com/redis/go-redis/v9"

	"fmt"
	"log"
	"net"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/filager"

	"google.golang.org/grpc"
)

const SERVICE_ID = 2

func main() {
	// config getting
	conf, err := config.GetConfig("../../config.yaml")
	if err != nil {
		log.Fatalf("Error with getting config: %v", err)
	}
	service_config, err := conf.GetServiceById(SERVICE_ID)
	if err != nil {
		log.Fatalf("Error with finding condfig: %v", err)
	}

	// connecting redis
	redis_client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port),
		Username: service_config.RedisConnect.User,
		Password: service_config.RedisConnect.Password, // No password set
		DB:       int(service_config.RedisConnect.DB),  // Use default DB
	})

	// starting server
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", service_config.Host, service_config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFileManagerServiceServer(grpcServer, server.NewServer(redis_client))

	fmt.Println(lis.Addr())
	grpcServer.Serve(lis)
}
