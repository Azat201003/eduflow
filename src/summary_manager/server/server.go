package server

import (

	// "user-service/server/db"

	"context"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/summager"
	"github.com/redis/go-redis/v9"
)

const DATA_FOLDER = "summaries/"
const FORMAT = ".md"

type summaryManagerServiceServer struct {
	pb.UnimplementedSummaryManagerServiceServer
	redis_client *redis.Client
}

func validateFilePath(filePath string) bool {
	return !(strings.Contains(filePath, ".") || strings.Contains(filePath, " "))
}

func (s *summaryManagerServiceServer) StartSending(context context.Context, request *pb.StartRequest) (*pb.StartResponse, error) {
	uuid := rand.N(67108864)
	val, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	os.Create(DATA_FOLDER + request.FilePath + FORMAT)

	err = s.redis_client.Set(context, fmt.Sprintf("sends[%v]", uuid), val, time.Second*time.Duration(2*request.FileSize)).Err()
	// TODO		add checking file in summary and clear before sending
	// !		add checking file in summary and clear before sending
	return &pb.StartResponse{Uuid: uint64(uuid)}, err
}

func (s *summaryManagerServiceServer) SendChunk(context context.Context, chunk *pb.WriteChunk) (*pb.WriteResponse, error) {
	r := s.redis_client.Get(context, fmt.Sprintf("sends[%v]", chunk.Uuid))
	if r.Err() != nil {
		return &pb.WriteResponse{Code: 0}, r.Err()
	}

	info := &pb.StartRequest{}
	bytes, err := r.Bytes()
	if err != nil {
		return &pb.WriteResponse{Code: -1}, err
	}

	err = json.Unmarshal(bytes, &info)

	f, err := os.OpenFile(DATA_FOLDER+info.FilePath+FORMAT, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return &pb.WriteResponse{Code: 1}, err
	}

	_, err = f.Write(chunk.Content)

	if err != nil {
		return &pb.WriteResponse{Code: 2}, err
	}

	return &pb.WriteResponse{Code: 3}, nil
}

func (s *summaryManagerServiceServer) CloseSending(context context.Context, request *pb.EndRequest) (*pb.EndResponse, error) {
	// TODO send request to summary to add file to database
	// include configuration
	// conf, err := config.GetConfig("../../config.yaml")
	// if err != nil {
	// 	log.Fatalf("Error with getting configuration %v", err.Error())
	// }

	// // summary_conf, err := conf.GetServiceById(1)

	// if err != nil {
	// 	log.Fatalf("Error with finding configuration %v", err.Error())
	// }

	r := s.redis_client.Get(context, fmt.Sprintf("sends[%v]", request.Uuid))
	if r.Err() != nil {
		return nil, r.Err()
	}

	info := &pb.StartRequest{}
	r.Scan(info)

	// // connect to summary service
	// summary_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", summary_conf.Host, summary_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	return nil, err
	// }

	// summary_client := summary.NewSummaryServiceClient(summary_conn)
	// _, err = summary_client.CreateSummary(context, &summary.Summary{Title: request.Title, Description: request.Description, FilePath: info.FilePath}, grpc.PeerCallOption{})

	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func NewServer(redis_client *redis.Client) pb.SummaryManagerServiceServer {
	server := new(summaryManagerServiceServer)
	server.redis_client = redis_client
	return server
}
