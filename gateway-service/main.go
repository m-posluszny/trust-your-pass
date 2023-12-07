package main

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"inf/gateway-service/configs"
	_ "log"
	"net/http"
)

type Password struct {
	id       int    `json:"id"`
	password string `json:"password"`
	strength int    `json:"strength"`
}

var PasswordColletion *mongo.Collection

func main() {

	configs.InitEnvConfigs()

	//Connect with mongo
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(configs.EnvList.MongoUri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	PasswordColletion = client.Database(configs.EnvList.DbName).Collection(configs.EnvList.CollectionName)

	router := gin.Default()
	router.GET("/api/v1/passwords", getPasswords)
	router.POST("/api/v1/passwords", postPassword)

	router.Run(configs.EnvList.ServerAddress)
}

func getPasswords(c *gin.Context) {
	filter := bson.D{{}}
	cursor, err := PasswordColletion.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []Password
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	//TODO fix parsing
	var outputJson [][]byte
	for _, result := range results {
		res, _ := json.Marshal(result)
		outputJson = append(outputJson, res)
	}
	res, _ := json.MarshalIndent(outputJson, "", "  ")
	c.IndentedJSON(http.StatusOK, res)
}

func postPassword(c *gin.Context) {
	var requestBody Password
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	cursor, _ := PasswordColletion.InsertOne(context.TODO(), requestBody)
	c.JSON(http.StatusCreated, cursor.InsertedID)
}
