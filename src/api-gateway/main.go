package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var user_client user.UserServiceClient

func authMiddlware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth_header := c.Request().Header["Authorization"]
		token := ""
		if len(auth_header) != 0 {
			token = strings.Split(auth_header[0], "Bearer ")[1]
		}
		var opts []grpc.CallOption
		authed_user, err := user_client.GetUserByToken(context.Background(), &user.Token{Token: token}, opts...)
		c.Set("is_user_setted", true)
		if err != nil {
			c.Set("is_user_setted", false)
		}
		c.Set("user", authed_user)

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}

func main() {
	// include configuration
	conf, err := config.GetConfig("../../config.yaml")
	if err != nil {
		log.Fatalf("Error with getting configuration %v", err.Error())
	}

	user_conf, err := conf.GetServiceById(0)

	if err != nil {
		log.Fatalf("Error with finding configuration %v", err.Error())
	}

	// connect to user service
	user_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", user_conf.Host, user_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error %v", err.Error())
	}
	user_client = user.NewUserServiceClient(user_conn)

	// echo endpoints
	e := echo.New()

	e.Use(authMiddlware)

	e.GET("/", func(c echo.Context) error {
		name := "world"
		if c.Get("is_user_setted").(bool) {
			var authed_user *user.User = c.Get("user").(*user.User)
			name = authed_user.GetUsername()
		}
		return c.String(http.StatusOK, fmt.Sprintf("Hello, %v!", name))
	})

	e.POST("auth/sign-up/", func(c echo.Context) error {
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
