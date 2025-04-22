package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// include configuration
	conf, err := config.GetConfig("../../config.yaml")
	if err != nil {
		log.Printf("error with getting configuration %v", err.Error())
	}

	user_conf, err := conf.GetServiceById(0)

	if err != nil {
		log.Printf("error with finding configuration %v", err.Error())
	}

	// connect to user service
	user_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", user_conf.Host, user_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error %v", err.Error())
	}
	user_client := user.NewUserServiceClient(user_conn)

	// echo endpoints
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	e.GET("auth/sign-up/", func(c echo.Context) error {
		var opts []grpc.CallOption
		username := c.FormValue("username")
		password := c.FormValue("password")
		token, err := user_client.Register(context.Background(), &user.Creditionals{Username: username, Password: password}, opts...)
		if err != nil {
			return err
		}
		return c.String(http.StatusCreated, token.Token)
	})

	e.GET("auth/sign-in/", func(c echo.Context) error {
		var opts []grpc.CallOption
		username := c.FormValue("username")
		password := c.FormValue("password")
		token, err := user_client.Login(context.Background(), &user.Creditionals{Username: username, Password: password}, opts...)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, token.Token)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
