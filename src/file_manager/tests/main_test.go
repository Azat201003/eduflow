package main

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	var opts []grpc.CallOption
	chunk_size := 12
	data := []byte(
		`# Something

It is **Something**, but I *don't know* what is it.

`)
	response, err := (*s.Client).StartSending(context.Background(), &filager.StartWriteRequest{FilePath: "new", ChunkSize: uint64(chunk_size), FileSize: uint64(len(data)), FileType: filager.FileType_DOCUMENT}, opts...)
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

// func TestReading(t *testing.T) {
// 	client, err := connectSummager()

// 	if err != nil {
// 		t.Error(err)
// 	}

// 	var opts []grpc.CallOption
// 	response, err := (*client).StartReading(context.Background(), &summager.StartReadRequest{FilePath: "new", ChunkSize: 8}, opts...)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	uuid := response.Uuid
// 	fmt.Println(uuid)
// 	data := ""
// 	for {
// 		resp, err := (*client).ReadChunk(context.Background(), &summager.ReadRequest{Uuid: int32(uuid)})
// 		// fmt.Println(err.Error())
// 		if err != nil {
// 			if s, _ := status.FromError(err); s.Message() == "EOF" {
// 				break
// 			}
// 			// fmt.Println(err.Error(), io.EOF)
// 			t.Error(err)
// 		}
// 		fmt.Println(err)
// 		data += string(resp.Content)
// 	}

// 	fmt.Println(uuid, "\n-----------------------\n", data, "\n-----------------------")
// }
