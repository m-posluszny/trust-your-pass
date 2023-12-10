package main

import (
	"context"
	"inf/gateway-service/configs"
	"inf/gateway-service/passwords"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//TODO set default value of strength to -1 or similar
//TODO use interfaces, controllers, interfacesImpl
//TODO put events to rabbitmq, add listener
//TODO tests
//TODO golang app docker compose

//TODO wątek dla rabbitmq

func main() {

	mongoClient := configs.GetMongoClient()
	defer mongoClient.Disconnect(context.TODO())

	//rmqConn := configs.GetConnection()
	//defer rmqConn.Close()

	router := gin.Default()
	passwords.SetupRoutes(router)
	router.Use(cors.Default())

	router.Run(configs.EnvList.ServerAddress)
}
