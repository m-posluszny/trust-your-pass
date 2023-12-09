package main

import (
	"context"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"inf/gateway-service/configs"
	"net/http"
)

//TODO set default value of strength to -1 or similar
//TODO use interfaces, controllers, interfacesImpl
//TODO put events to rabbitmq, add listener
//TODO tests

type PasswordDto struct {
	//Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Password string `json:"password" bson:"password"`
	Strength int    `json:"strength" bson:"strength"`
}

var PasswordCollection *mongo.Collection
var MongoClient *mongo.Client

var RmqConnection *amqp.Connection
var RmqChannel *amqp.Channel

func main() {

	MongoClient, PasswordCollection = configs.InitMongoConnection()
	defer MongoClient.Disconnect(context.TODO())

	RmqConnection, RmqChannel = configs.InitRabbitMQConnection()

	router := gin.Default()
	router.GET("/api/v1/passwords", getPasswords)
	router.GET("/api/v1/passwords/:password", getPasswordStrength)
	router.GET("/api/v1/passwords/:password/poll", pollResult)
	router.POST("/api/v1/passwords", postPassword)

	router.Run(configs.EnvList.ServerAddress)
}

func getPasswords(c *gin.Context) {
	cursor, err := PasswordCollection.Find(c, bson.M{"strength": bson.M{"$ne": -1}})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	var results []PasswordDto
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, results)
}

func getPasswordStrength(c *gin.Context) {
	password := c.Param("password")

	isValid, preconditions := validateInput(password)
	if !isValid {
		c.JSON(http.StatusBadRequest, preconditions)
		return
	}

	var queryResult PasswordDto
	err := PasswordCollection.FindOne(context.TODO(), bson.D{{"password", password}}).Decode(&queryResult)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, queryResult.Strength)
}

func postPassword(c *gin.Context) {
	requestBody := PasswordDto{}

	//bind req body
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//validate if string consists of all spaces, etc
	isValid, preconditions := validateInput(requestBody.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, preconditions)
		return
	}

	filter := bson.D{{"password", requestBody.Password}}
	update := bson.D{{"$set", bson.D{{"password", requestBody.Password}, {"strength", requestBody.Strength}}}}
	opts := options.Update().SetUpsert(true)
	result, err := PasswordCollection.UpdateOne(context.TODO(), filter, update, opts)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, result.UpsertedID)
}

func validateInput(password string) (bool, string) {
	var preconditions []string

	if len(password) <= 3 {
		preconditions = append(preconditions, "Password should contain at least 4 signs")
	}
	if stringUtils.IsBlank(password) {
		preconditions = append(preconditions, "Provided string is empty")
	}
	if stringUtils.IsWhitespace(password) {
		preconditions = append(preconditions, "Provided string consists of whitespaces only")
	}
	if stringUtils.ContainsAny(password, "\\0") {
		preconditions = append(preconditions, "Provided string contains not allowed symbols")
	}

	if len(preconditions) == 0 {
		return true, ""
	} else {
		return false, stringUtils.Join(preconditions, "\n")
	}
}

// might be unused if model will return results at fast manner
func pollResult(c *gin.Context) {
	password := c.Param("password")
	password += ""
}
