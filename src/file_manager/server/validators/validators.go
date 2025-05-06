package validators

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	pb "github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/Azat201003/eduflow_service_api/gen/go/summary"
)

type Validator interface {
	Validate(interface{}) error
}

type MyValidator struct {
	SummaryClient summary.SummaryServiceClient
}

func (v *MyValidator) Validate(obj interface{}) error {
	switch t := obj.(type) {
	case *pb.StartWriteRequest:
		return v.ValidateStartWriteRequest(t)
	case *pb.WriteChunk:
		return v.ValidateWriteChunk(t)
	default:
		return errors.New("Validator error: unknow type")
	}
}

func (v *MyValidator) ValidateStartWriteRequest(obj *pb.StartWriteRequest) error {
	if err := v.ValidateFilePath(obj.FilePath); err != nil {
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

func (v *MyValidator) ValidateFilePath(obj string) error {
	if obj == "" {
		return errors.New("FilePath is empty")
	}
	if strings.Contains(obj, ".") || strings.Contains(obj, " ") || strings.Contains(obj, "\n") || obj[0] == '/' {
		return errors.New("FilePath contains invalid characters")
	}

	stream, err := (v.SummaryClient).GetFilteredSummaries(context.Background(), &summary.FilterRequest{
		Filter: &summary.Summary{FilePath: obj},
		Page: &summary.Page{
			Size:   1,
			Number: 1,
		},
	})
	if err == nil {
		if data, err := stream.Recv(); err != io.EOF {
			log.Println(data, err)
			return errors.New("This filepath already exist.")
		}
		stream.CloseSend()
	}
	return nil
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

func (v *MyValidator) ValidateWriteChunk(obj *pb.WriteChunk) error {
	if len(obj.Content) == 0 {
		return errors.New("asfd") // TODO create redis manager
	}
	return nil
}
