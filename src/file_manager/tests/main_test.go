package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand/v2"
	"strconv"
	"testing"
	"time"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type ClientTestSuite struct {
	suite.Suite
	Client *filager.FileManagerServiceClient
}

func TestClientSuite(t *testing.T) {
	t.Helper()
	t.Parallel()

	conf, err := config.GetConfig("../../config.yaml")
	assert.NoError(t, err)
	filager_conf, err := conf.GetServiceById(2)
	assert.NoError(t, err)
	filager_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", filager_conf.Host, filager_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	filager_client := filager.NewFileManagerServiceClient(filager_conn)

	s := ClientTestSuite{
		Client: &filager_client,
	}
	suite.Run(t, &s)
}

func (s *ClientTestSuite) TestWriting() {
	chunk_size := 12
	data := []byte(
		`# Something

It is **Something**, but I *don't know* what is it.

`)
	id := strconv.Itoa(rand.Int())
	response, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{FilePath: "test-" + id, ChunkSize: uint64(chunk_size), FileSize: uint64(len(data)), FileType: filager.FileType_DOCUMENT})
	s.NoError(err)
	log.Println(id)
	uuid := response.Uuid
	for i := 0; ; i++ {
		if i*chunk_size >= len(data) {
			break
		}
		chunk := make([]byte, chunk_size)
		l := i * chunk_size
		r := int(math.Min(float64((i+1)*chunk_size), float64(len(data)-1)))
		for j := l; j < r; j++ {
			chunk[j-l] = data[j]
		}
		for j := r; j < chunk_size; j++ {
			chunk[j-l] = '\x00'
		}

		_, err = (*s.Client).SendChunk(context.Background(), &filager.WriteChunk{Uuid: uuid, Content: chunk})
		s.NoError(err)
	}
	r, err := (*s.Client).CloseSending(context.Background(), &filager.EndRequest{Uuid: uuid, Title: "New", Description: "New file", AuthorId: 2})
	s.NoError(err)
	fmt.Println(r)
}

func (s *ClientTestSuite) TestWriteExistError() {
	_, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{FilePath: "test", ChunkSize: 8, FileSize: 32, FileType: filager.FileType_DOCUMENT})
	s.ErrorContains(err, "exist")
}

func sliceEqual(a, b []byte) bool {
	if len(a) > len(b) {
		a, b = b, a
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (s *ClientTestSuite) TestReading() {
	chunk_size := 18
	response, err := (*s.Client).StartReading(context.Background(), &filager.StartReadRequest{
		ChunkSize: uint64(chunk_size),
		FilePath:  "test",
		FileType:  filager.FileType_DOCUMENT,
	})
	s.NoError(err)
	uuid := response.Uuid

	var loaded_data []byte
	for i := 0; ; i++ {
		chunk, err := (*s.Client).ReadChunk(context.Background(), &filager.ReadRequest{Uuid: uuid, ChunkNumber: uint64(i)})
		if s, _ := status.FromError(err); s.Message() == "EOF" {
			break
		}
		s.NoError(err)
		loaded_data = append(loaded_data, chunk.Content...)
	}

	fmt.Println(string(loaded_data))
}

func (s *ClientTestSuite) TestEmptyFilePath() {
	_, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{
		ChunkSize: 4,
		FileSize:  16,
		FileType:  filager.FileType_DOCUMENT,
	})
	s.Error(err)
}

func (s *ClientTestSuite) TestEmptyFileSize() {
	_, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{
		ChunkSize: 4,
		FilePath:  "abeme_52",
		FileType:  filager.FileType_DOCUMENT,
	})
	s.Error(err)
}

func (s *ClientTestSuite) TestEmptyChunkSize() {
	_, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{
		FilePath: "abeme_1234",
		FileSize: 16,
		FileType: filager.FileType_DOCUMENT,
	})
	s.Error(err)
}

func (s *ClientTestSuite) TestTimeOut() {
	resp, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{
		FilePath:  "abeme_1234",
		FileSize:  8,
		ChunkSize: 4,
		FileType:  filager.FileType_DOCUMENT,
	})
	s.NoError(err)
	uuid := resp.Uuid
	time.Sleep(time.Second * time.Duration(9))
	_, err = (*s.Client).SendChunk(context.Background(), &filager.WriteChunk{Uuid: uuid, Content: []byte("# Ab")})
	s.Error(err)
}
