package main

import (
	"context"
	"inf/gateway-service/configs"
	"inf/gateway-service/passwords"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//TODO tests

func main() {

	mongoClient := configs.GetMongoClient()
	defer mongoClient.Disconnect(context.TODO())

	rmqConn := configs.GetConnection()
	defer rmqConn.Close()
	go passwords.ConsumeMessages()

	router := gin.Default()
	passwords.SetupRoutes(router)
	router.Use(cors.Default())

	router.Run(configs.EnvList.ServerAddress)

}
