package validators

import (
	"context"
	"errors"
	redis_manager "filager-service/server/redis"
	"fmt"
	"io"
	"log"
	"strings"

	pb "github.com/Azat201003/eduflow/backend/libs/gen/go/filager"
	"github.com/Azat201003/eduflow/backend/libs/gen/go/summary"
)

type Validator interface {
	Validate(interface{}) error
}

type MyValidator struct {
	SummaryClient *summary.SummaryServiceClient
	RedisManager  *redis_manager.RedisManager
}

func NewMyValidator(summaryClient *summary.SummaryServiceClient, redisManager *redis_manager.RedisManager) *MyValidator {
	return &MyValidator{
		SummaryClient: summaryClient,
		RedisManager:  redisManager,
	}
}

func (v *MyValidator) Validate(obj interface{}) error {
	switch t := obj.(type) {
	case *pb.StartWriteRequest:
		return v.ValidateStartWriteRequest(t)
	case *pb.WriteChunk:
		return v.ValidateWriteChunk(t)
	case *pb.EndRequest:
		return v.ValidateEndRequest(t)
	case *pb.StartReadRequest:
		return v.ValidateStartReadRequest(t)
	case *pb.ReadRequest:
		return v.ValidateReadRequest(t)
	default:
		return errors.New("Validator error: unknown type")
	}
}

func (c *MyValidator) ValidateEndRequest(obj *pb.EndRequest) error {
	if obj.AuthorId == 0 {
		return errors.New("No author setted")
	}
	return nil
}

func (v *MyValidator) ValidateWriteChunk(obj *pb.WriteChunk) error {
	if obj == nil {
		return errors.New("WriteChunk is nil")
	}
	session, err := v.RedisManager.GetSendingSession(obj.Uuid)
	if err != nil {
		return nil
	}

	if session.ChunkSize != uint64(len(obj.Content)) {
		return errors.New(fmt.Sprintf("Unmatching size of chunk: should be %v, but got %v", session.ChunkSize, len(obj.Content)))
	}

	return nil
}

func (v *MyValidator) ValidateStartReadRequest(obj *pb.StartReadRequest) error {
	if err := v.ValidateReadFilePath(obj.FilePath); err != nil {
		return err
	}
	if err := v.ValidateChunkSize(obj.ChunkSize); err != nil {
		return err
	}
	if obj.FileType == pb.FileType_UNDEFIGNED {
		return errors.New("FileType is unspecified")
	}
	return nil
}

func (v *MyValidator) ValidateReadFilePath(obj string) error {
	stream, err := (*v.SummaryClient).GetFilteredSummaries(context.Background(), &summary.FilterRequest{
		Filter: &summary.Summary{FilePath: obj},
		Page: &summary.Page{
			Size:   1,
			Number: 1,
		},
	})
	log.Println(err, obj)
	if err == nil {
		if data, err := stream.Recv(); err != io.EOF {
			log.Println(data, err)
			return nil
		}
		stream.CloseSend()
		return errors.New("File path not exist")
	}
	return errors.New("Cannot connect to summary service")
}

func (v *MyValidator) ValidateReadRequest(obj *pb.ReadRequest) error {
	return nil
}

func (v *MyValidator) ValidateStartWriteRequest(obj *pb.StartWriteRequest) error {
	if err := v.ValidateWriteFilePath(obj.FilePath); err != nil {
		return err
	}
	if err := v.ValidateChunkSize(obj.ChunkSize); err != nil {
		return err
	}
	if err := v.ValidateFileSize(obj.FileSize); err != nil {
		return err
	}
	return nil
}

func (v *MyValidator) ValidateWriteFilePath(obj string) error {
	if obj == "" {
		return errors.New("FilePath is empty")
	}
	if strings.Contains(obj, ".") || strings.Contains(obj, " ") || strings.Contains(obj, "\n") || obj[0] == '/' {
		return errors.New("FilePath contains invalid characters")
	}

	stream, err := (*v.SummaryClient).GetFilteredSummaries(context.Background(), &summary.FilterRequest{
		Filter: &summary.Summary{FilePath: obj},
		Page: &summary.Page{
			Size:   1,
			Number: 1,
		},
	})
	log.Println(err, obj)
	if err == nil {
		if data, err := stream.Recv(); err != io.EOF {
			log.Println(data, err)
			return errors.New("This filepath already exist.")
		}
		stream.CloseSend()
		return nil
	}
	return errors.New("Cannot connect to summary service")
}

func (v *MyValidator) ValidateChunkSize(obj uint64) error {
	if obj == 0 {
		return errors.New("ChunkSize is 0")
	}
	return nil
}

func (v *MyValidator) ValidateFileSize(obj uint64) error {
	if obj == 0 {
		return errors.New("FileSize is 0")
	}
	return nil
}
