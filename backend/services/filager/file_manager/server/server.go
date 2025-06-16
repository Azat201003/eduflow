package server

import (

	// "user-service/server/db"

	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	redis_manager "filager-service/server/redis"
	validators "filager-service/server/validators"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/Azat201003/eduflow_service_api/gen/go/summary"
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
	summary_client *summary.SummaryServiceClient
	redisManager   *redis_manager.RedisManager
	validator      validators.Validator
}

func (s *fileManagerServiceServer) StartSending(context context.Context, request *pb.StartWriteRequest) (*pb.StartResponse, error) {
	// validate filepath
	if err := s.validator.Validate(request); err != nil {
		return nil, err
	}

	// clear and creating filepath
	_, err := os.Create(DATA_FOLDER + request.FilePath + formatById(request.FileType))
	if err != nil {
		return nil, err
	}

	// creating session
	uuid, err := s.redisManager.CreateSendingSession(request)

	return &pb.StartResponse{Uuid: uuid}, err
}

func (s *fileManagerServiceServer) SendChunk(context context.Context, chunk *pb.WriteChunk) (*pb.WriteResponse, error) {
	if err := s.validator.Validate(chunk); err != nil {
		return nil, err
	}

	session, err := s.redisManager.GetSendingSession(chunk.Uuid)
	if err != nil {
		return &pb.WriteResponse{Code: -1}, err
	}

	f, err := os.OpenFile(DATA_FOLDER+session.FilePath+formatById(session.FileType), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
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
	if err := s.validator.Validate(request); err != nil {
		return nil, err
	}

	session, err := s.redisManager.GetSendingSession(request.Uuid)
	if err != nil {
		return nil, err
	}

	_, err = (*s.summary_client).CreateSummary(context, &summary.Summary{Title: request.Title, Description: request.Description, FilePath: session.FilePath, AuthorId: &summary.Id{Id: request.AuthorId}})

	if err != nil {
		return nil, err
	}

	s.redisManager.CloseSendingSession(request.Uuid)

	return &pb.EndResponse{Code: 0}, nil
}

func (s *fileManagerServiceServer) StartReading(context context.Context, request *pb.StartReadRequest) (*pb.StartResponse, error) {
	if err := s.validator.Validate(request); err != nil {
		return nil, err
	}

	f, err := os.Open(DATA_FOLDER + request.FilePath + formatById(request.FileType))
	if err != nil {
		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	uuid, err := s.redisManager.CreateReadingSession(&redis_manager.ReadingSession{
		FilePath:  request.FilePath,
		FileType:  request.FileType,
		FileSize:  uint64(fi.Size()),
		ChunkSize: uint16(request.ChunkSize), // TODO сделать одного формата (uint32/uint64/int32/int64)
	})

	return &pb.StartResponse{Uuid: uint64(uuid)}, err
}

func (s *fileManagerServiceServer) ReadChunk(context context.Context, request *pb.ReadRequest) (*pb.GetChunk, error) {
	if err := s.validator.Validate(request); err != nil {
		return nil, err
	}

	session, err := s.redisManager.GetReadingSession(request.Uuid)
	log.Println(fmt.Sprintf("reads[%v]", request.Uuid))

	f, err := os.Open(DATA_FOLDER + session.FilePath + formatById(session.FileType))
	if err != nil {
		return nil, err
	}
	if session.FileSize <= uint64(request.ChunkNumber)*uint64(session.ChunkSize) {
		return nil, errors.New("EOF")
	}
	data := make([]byte, session.ChunkSize)
	_, err = f.ReadAt(data, int64(request.ChunkNumber*uint64(session.ChunkSize)))

	if err != nil && err != io.EOF {
		return nil, err
	}

	return &pb.GetChunk{Content: data}, nil

}

func NewServer(redisManager *redis_manager.RedisManager, summaryClient *summary.SummaryServiceClient, validator validators.Validator) pb.FileManagerServiceServer {
	server := new(fileManagerServiceServer)
	server.redisManager = redisManager
	server.summary_client = summaryClient
	server.validator = validator
	return server
}
