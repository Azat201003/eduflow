package main

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"strconv"
	"testing"

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

	conf, err := config.GetConfig("../../../config.yaml")
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
	response, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{FilePath: "test-" + strconv.Itoa(rand.Int()), ChunkSize: uint64(chunk_size), FileSize: uint64(len(data)), FileType: filager.FileType_DOCUMENT})
	s.NoError(err)
	uuid := response.Uuid
	for i := 0; ; i++ {
		if i*chunk_size >= len(data) {
			break
		}
		_, err = (*s.Client).SendChunk(context.Background(), &filager.WriteChunk{Uuid: uuid, Content: data[i*chunk_size : int(math.Min(float64((i+1)*chunk_size), float64(len(data)-1)))]})
		s.NoError(err)
	}
	r, err := (*s.Client).CloseSending(context.Background(), &filager.EndRequest{Uuid: uuid, Title: "New", Description: "New file", AuthorId: 2})
	s.NoError(err)
	fmt.Println(r)
}

func (s *ClientTestSuite) TestWriteExistError() {
	_, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{FilePath: "new"})
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
		FilePath:  "new",
		FileType:  filager.FileType_DOCUMENT,
	})
	s.NoError(err)
	uuid := response.Uuid

	var loaded_data []byte
	for {
		chunk, err := (*s.Client).ReadChunk(context.Background(), &filager.ReadRequest{Uuid: int32(uuid)})
		if s, _ := status.FromError(err); s.Message() == "EOF" {
			break
		}
		s.NoError(err)
		loaded_data = append(loaded_data, chunk.Content...)
	}

	fmt.Println(string(loaded_data))
}
