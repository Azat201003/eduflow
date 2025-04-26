package main

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var user_client user.UserServiceClient

func connectUser() (user.UserServiceClient, error) {
	// include configuration
	conf, err := config.GetConfig("../../../config.yaml")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with getting configuration %v", err.Error()))
	}

	user_conf, err := conf.GetServiceById(0)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("error with finding configuration %v", err.Error()))
	}

	// connect to user service
	user_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", user_conf.Host, user_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error %v", err.Error()))
	}
	user_client = user.NewUserServiceClient(user_conn)
	return user_client, nil
}

func TestLoggingIn(t *testing.T) {
	user_client, err := connectUser()
	if err != nil {
		t.Fatal(err)
	}
	var opts []grpc.CallOption
	token, err := user_client.Login(context.Background(), &user.Creditionals{Username: "Coolman", Password: "1234"}, opts...)
	if err != nil {
		t.Fatal(err)
	}
	if token.Token != "anqrfsNqOu" {
		fmt.Println("Getted token: ", token.Token)
		t.Fatal("Token is not valid")
	}
}

func TestGettingByToken(t *testing.T) {
	user_client, err := connectUser()
	if err != nil {
		t.Fatal(err)
	}
	var opts []grpc.CallOption
	user_obj, err := user_client.GetUserByToken(context.Background(), &user.Token{Token: "anqrfsNqOu"}, opts...)
	if err != nil {
		t.Fatal(err)
	}
	if user_obj.Username != "Coolman" {
		fmt.Println("Getted user: ", user_obj)
		t.Fatal("User is not valid")
	}
}

func TestGettingById(t *testing.T) {
	user_client, err := connectUser()
	if err != nil {
		t.Fatal(err)
	}
	var opts []grpc.CallOption
	user_obj, err := user_client.GetUserById(context.Background(), &user.Id{Id: 3}, opts...)
	if err != nil {
		t.Fatal(err)
	}
	if user_obj.Username != "Coolman" {
		fmt.Println("Getted user: ", user_obj)
		t.Fatal("User is not valid")
	}
}
