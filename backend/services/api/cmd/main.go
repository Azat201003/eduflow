package main

import (
	"api-gateway/middlwares"
	"api-gateway/routes"
	"fmt"
	"log"

	"github.com/Azat201003/eduflow_service_api/config"
	"github.com/Azat201003/eduflow_service_api/gen/go/filager"
	"github.com/Azat201003/eduflow_service_api/gen/go/summary"
	"github.com/Azat201003/eduflow_service_api/gen/go/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// func corsMiddlware() {

// }

func main() {
	// include configuration
	conf, err := config.GetConfig("../../config.yaml")
	if err != nil {
		log.Fatalf("Error with getting configuration %v", err.Error())
	}

	user_conf, err := conf.GetServiceById(0)
	summary_conf, err := conf.GetServiceById(1)
	filager_conf, err := conf.GetServiceById(2)

	if err != nil {
		log.Fatalf("Error with finding configuration %v", err.Error())
	}

	// connect to user service
	user_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", user_conf.Host, user_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error %v", err.Error())
	}
	routes.UserClient = user.NewUserServiceClient(user_conn)

	// connect to summary service
	summary_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", summary_conf.Host, summary_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error %v", err.Error())
	}
	routes.SummaryClient = summary.NewSummaryServiceClient(summary_conn)

	// connect to summary service
	filager_conn, err := grpc.NewClient(fmt.Sprintf("%v:%v", filager_conf.Host, filager_conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error %v", err.Error())
	}
	routes.FilagerClient = filager.NewFileManagerServiceClient(filager_conn)

	// echo endpoints
	e := echo.New()

	e.Use(middlwares.AuthMiddlware)

	e.GET("/", routes.MainRoute)
	e.POST("auth/sign-up/", routes.SignUp)
	e.POST("auth/sign-in/", routes.SignIn)
	e.GET("my-user/", routes.MyUser)
	e.OPTIONS("my-user/", func(c echo.Context) error {
		c.Response().Header().Add("Access-Control-Allow-Origin", "*")
		c.Response().Header().Add("Access-Control-Allow-Methods", "GET")
		c.Response().Header().Add("Access-Control-Allow-Headers", "Authorization")
		return nil
	})
	e.GET("users/:id/", routes.GetUserById)
	e.GET("summaries/list/", routes.SummaryList)
	e.GET("summaries/:id/", routes.GetSummaryById)
	e.POST("files/start-sending/", routes.StartFileSending)
	e.POST("files/send-chunk/", routes.SendChunk)
	e.POST("files/close-sending/", routes.CloseFileSending)

	e.Logger.Fatal(e.Start(":8080"))
}
