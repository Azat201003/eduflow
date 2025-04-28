package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/go/summager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectSummager() (*summager.SummaryManagerServiceClient, error) {
	conf, err := config.GetConfig("../../../config.yaml")
	if err != nil {
		return nil, err
	}

	summager_conf, err := conf.GetServiceById(2)

	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", summager_conf.Host, summager_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := summager.NewSummaryManagerServiceClient(conn)
	return &client, nil
}

func TestWriting(t *testing.T) {
	client, err := connectSummager()

	if err != nil {
		t.Error(err)
	}

	var opts []grpc.CallOption
	response, err := (*client).StartSending(context.Background(), &summager.StartRequest{FilePath: "abeme", ChunkSize: 8, FileSize: 8}, opts...)
	if err != nil {
		t.Error(err)
	}
	uuid := response.Uuid

	_, err = (*client).SendChunk(context.Background(), &summager.WriteChunk{Uuid: uuid, Content: []byte("# Abeme\n")})
	if err != nil {
		t.Error(err)
	}
	r, err := (*client).CloseSending(context.Background(), &summager.EndRequest{Uuid: uuid, Title: "Abeme", Description: "Abeme rules the world"})

	if err != nil {
		t.Error(err)
	}

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

func TestReading(t *testing.T) {
	// client, err := connectSummager()

	// if err != nil {
	// 	t.Error(err)
	// }

	// var opts []grpc.CallOption
	// stream, err := (*client).StartReading(context.Background(), &summager.StartRequest{FilePath: "соли", ChunkSize: 64}, opts...)
	// if err != nil {
	// 	t.Error(err)
	// }

	// block := make([]byte, 64)

	// chunk, err := stream.Recv()
	// block = chunk.Content
	// should_be := make([]byte, 64)
	// should_be = []byte("# Соли")
	// fmt.Println(string(block))
	// if !sliceEqual(block, should_be) {
	// 	t.Error("Invalid content error")
	// }

	// if _, err := stream.Recv(); err != io.EOF {
	// 	t.Error(err)
	// }

	// err = stream.CloseSend()
	// if err != nil {
	// 	t.Error(err)
	// }
}
