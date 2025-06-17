package main

import (
	"filager/internal/server"
	"filager/internal/validators"
	redis_manager "filager/internal/redis"

	"fmt"
	"log"
	"net"

	"github.com/redis/go-redis/v9"

	config "github.com/Azat201003/eduflow/backend/libs/config"
	pb "github.com/Azat201003/eduflow/backend/libs/gen/go/filager"
	"github.com/Azat201003/eduflow/backend/libs/gen/go/summary"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const SERVICE_ID = 2

func main() {
	// # Config getting
	conf, err := config.GetConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Error with getting config: %v", err)
	}
	service_config, err := conf.GetServiceById(SERVICE_ID)
	if err != nil {
		log.Fatalf("Error with finding condfig: %v", err)
	}

	// ## Summary
	if err != nil {
		log.Fatalf("Error with getting configuration %v", err.Error())
	}

	summary_conf, err := conf.GetServiceById(1)

	if err != nil {
		log.Fatalf("Error with finding configuration %v", err.Error())
	}

	// # Connecting redis
	redis_client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", conf.Redis.Host, conf.Redis.Port),
		Username: service_config.RedisConnect.User,
		Password: service_config.RedisConnect.Password, // No password set
		DB:       int(service_config.RedisConnect.DB),  // Use default DB
	})

	// # Connecting summary service
	summary_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", summary_conf.Host, summary_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error with connecting to summary service: ", err)
	}

	summary_client := summary.NewSummaryServiceClient(summary_conn)

	// # Starting server
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%d", service_config.Host, service_config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFileManagerServiceServer(grpcServer, server.NewServer(redis_manager.NewRedisManager(redis_client), &summary_client, validators.NewMyValidator(&summary_client, redis_manager.NewRedisManager(redis_client))))

	fmt.Println(lis.Addr())
	grpcServer.Serve(lis)
}
