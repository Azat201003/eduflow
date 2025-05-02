package server

import (

	// "user-service/server/db"

	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"strings"
	"time"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/Azat201003/eduflow_service_api/gen/go/summary"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

const DATA_FOLDER = "files/"

func formatById(id pb.FileType) string {
	return [...]string{
		pb.FileType_DOCUMENT: ".md",
		pb.FileType_IMAGE:    ".png",
	}[id]
}

type fileManagerServiceServer struct {
	pb.UnimplementedFileManagerServiceServer
	redis_client   *redis.Client
	summary_client *summary.SummaryServiceClient
}

type ReadingStatus struct {
	FilePath     string
	ChunksReaded uint16
	ChunkSize    uint16
	FileSize     uint64
	FileType     pb.FileType
}

func validateFilePath(filePath string) bool {
	return !(strings.Contains(filePath, ".") || strings.Contains(filePath, " "))
}

func (s *fileManagerServiceServer) StartSending(context context.Context, request *pb.StartWriteRequest) (*pb.StartResponse, error) {
	uuid := rand.NewPCG(uint64(time.Now().Nanosecond()), uint64(time.Now().Second())).Uint64()
	defer log.Println(uuid, uint64(time.Now().Nanosecond()), uint64(time.Now().Second()))
	val, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	_, err = os.Create(DATA_FOLDER + request.FilePath + formatById(request.FileType))

	if err != nil {
		return nil, err
	}

	err = s.redis_client.Set(context, fmt.Sprintf("sends[%v]", uuid), val, time.Second*time.Duration(2*request.FileSize)).Err()
	// TODO		add checking file in summary and clear before sending
	// !		add checking file in summary and clear before sending
	return &pb.StartResponse{Uuid: uuid}, err
}

func (s *fileManagerServiceServer) SendChunk(context context.Context, chunk *pb.WriteChunk) (*pb.WriteResponse, error) {
	r := s.redis_client.Get(context, fmt.Sprintf("sends[%v]", chunk.Uuid))
	if r.Err() != nil {
		return &pb.WriteResponse{Code: 0}, r.Err()
	}

	info := &pb.StartWriteRequest{}
	bytes, err := r.Bytes()
	if err != nil {
		return &pb.WriteResponse{Code: -1}, err
	}

	err = json.Unmarshal(bytes, info)

	f, err := os.OpenFile(DATA_FOLDER+info.FilePath+formatById(info.FileType), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return &pb.WriteResponse{Code: 1}, err
	}

	_, err = f.Write(chunk.Content)

	if err != nil {
		return &pb.WriteResponse{Code: 2}, err
	}

	return &pb.WriteResponse{Code: 3}, nil
}

func (s *fileManagerServiceServer) CloseSending(context context.Context, request *pb.EndRequest) (*pb.EndResponse, error) {
	// TODO send request to summary to add file to database
	// include configuration
	// conf, err := config.GetConfig("../../config.yaml")
	// if err != nil {
	// 	log.Fatalf("Error with getting configuration %v", err.Error())
	// }

	// summary_conf, err := conf.GetServiceById(1)

	// if err != nil {
	// 	log.Fatalf("Error with finding configuration %v", err.Error())
	// }

	r := s.redis_client.Get(context, fmt.Sprintf("sends[%v]", request.Uuid))
	if r.Err() != nil {
		return nil, r.Err()
	}

	info := &pb.StartWriteRequest{}
	bytes, err := r.Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, info)

	// connect to summary service
	// summary_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", summary_conf.Host, summary_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	return nil, err
	// }

	// summary_client := summary.NewSummaryServiceClient(summary_conn)
	_, err = (*s.summary_client).CreateSummary(context, &summary.Summary{Title: request.Title, Description: request.Description, FilePath: info.FilePath, AuthorId: &summary.Id{Id: request.AuthorId}}, []grpc.CallOption{}...)

	if err != nil {
		return nil, err
	}

	return &pb.EndResponse{Code: 0}, nil
}

func (s *fileManagerServiceServer) StartReading(context context.Context, request *pb.StartReadRequest) (*pb.StartResponse, error) {
	f, err := os.Open(DATA_FOLDER + request.FilePath + formatById(request.FileType))

	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	uuid := rand.N(536870912)
	val, err := json.Marshal(&ReadingStatus{
		ChunkSize:    uint16(request.ChunkSize),
		ChunksReaded: 0,
		FilePath:     request.FilePath,
		FileSize:     uint64(fi.Size()),
		FileType:     request.FileType,
	})
	if err != nil {
		return nil, err
	}

	err = s.redis_client.Set(context, fmt.Sprintf("reads[%v]", uuid), val, time.Second*time.Duration(fi.Size())).Err()
	log.Println(time.Second * time.Duration(2*fi.Size()))
	// TODO		add checking file in summary and clear before sending
	// !		add checking file in summary and clear before sending
	log.Println(string(val), ReadingStatus{ChunkSize: uint16(request.ChunkSize), ChunksReaded: 0, FilePath: request.FilePath, FileSize: uint64(fi.Size())})
	return &pb.StartResponse{Uuid: uint64(uuid)}, err
}

func (s *fileManagerServiceServer) ReadChunk(context context.Context, request *pb.ReadRequest) (*pb.GetChunk, error) {
	r := s.redis_client.Get(context, fmt.Sprintf("reads[%v]", request.Uuid))
	log.Println(fmt.Sprintf("reads[%v]", request.Uuid))
	if r.Err() != nil {
		return nil, r.Err()
	}

	info := &ReadingStatus{}
	bytes, err := r.Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, info)
	fmt.Println(string(bytes))
	become, err := json.Marshal(ReadingStatus{FilePath: info.FilePath, ChunksReaded: info.ChunksReaded + 1, ChunkSize: info.ChunkSize, FileSize: info.FileSize})

	if err != nil {
		return nil, err
	}

	f, err := os.Open(DATA_FOLDER + info.FilePath + formatById(info.FileType))
	if err != nil {
		return nil, err
	}

	if info.FileSize <= uint64(info.ChunksReaded)*uint64(info.ChunkSize) {
		fmt.Println(*info, &pb.GetChunk{Content: []byte("abeme")}, io.EOF)
		return nil, errors.New("EOF")
	} else {
		data := make([]byte, info.ChunkSize)
		_, err = f.ReadAt(data, int64(info.ChunksReaded*uint16(info.ChunkSize)))

		if err != nil {
			return nil, err
		}
		err = s.redis_client.Set(context, fmt.Sprintf("reads[%v]", request.Uuid), become, time.Second*time.Duration(info.FileSize-uint64(info.ChunkSize)*uint64(info.ChunksReaded))/2).Err()
		if err != nil {
			return nil, err
		}
		return &pb.GetChunk{Content: data}, nil
	}
}

func NewServer(redis_client *redis.Client) pb.FileManagerServiceServer {
	server := new(fileManagerServiceServer)
	server.redis_client = redis_client
	return server
}
