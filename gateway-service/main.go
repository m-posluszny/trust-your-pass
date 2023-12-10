package main

import (
	"context"
	"github.com/agrison/go-commons-lang/stringUtils"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"inf/gateway-service/configs"
	"net/http"
)

//TODO set default value of strength to -1 or similar
//TODO use interfaces, controllers, interfacesImpl
//TODO put events to rabbitmq, add listener
//TODO tests
//TODO golang app docker compose

//TODO wątek dla rabbitmq

type PasswordDto struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Password    string             `json:"password,omitempty" bson:"password,omitempty"`
	Strength    int                `json:"strength" bson:"strength"`
	IsProcessed bool               `json:"isProcessed" bson:"isProcessed"`
}

type PreconditionDto struct {
	ConditionName string `json:"condition"`
	IsSatisfied   bool   `json:"isSatisfied"`
}

func NewPreconditionDto(condition string) *PreconditionDto {
	return &PreconditionDto{condition, false}
}

type PostResponseDto struct {
	Id            interface{}     `json:"id" bson:"id"`
	Preconditions map[string]bool `json:"preconditions" bson:"preconditions"`
	IsProcessed   bool            `json:"isProcessed" bson:"isProcessed"`
}

func NewPostResponseDto(preconditions map[string]bool) *PostResponseDto {
	return &PostResponseDto{nil, preconditions, false}
}

var PasswordCollection *mongo.Collection
var MongoClient *mongo.Client

var RmqConnection *amqp.Connection
var RmqChannel *amqp.Channel

func main() {

	MongoClient, PasswordCollection = configs.InitMongoConnection()
	defer MongoClient.Disconnect(context.TODO())

	RmqConnection, RmqChannel = configs.InitRabbitMQConnection()
	defer RmqConnection.Close()

	router := gin.Default()
	router.GET("/api/v1/passwords", getAll)
	router.GET("/api/v1/passwords/:password", getByPassword)
	router.POST("/api/v1/passwords", insert)

	//TODO should be done by listener instead?
	router.PATCH("/api/v1/passwords/:password/update", updateStrength)

	router.Run(configs.EnvList.ServerAddress)
}

func getAll(c *gin.Context) {
	//get all records that were processed
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

func getByPassword(c *gin.Context) {
	password := c.Param("password")

	preconditions := validateInput(password)
	isValid := arePreconditionsSatisfied(preconditions)
	//why do I use 400 in that case: https://stackoverflow.com/questions/3290182/which-status-code-should-i-use-for-failed-validations-or-invalid-duplicates
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
	c.IndentedJSON(http.StatusOK, PasswordDto{
		Id:          queryResult.Id,
		Strength:    queryResult.Strength,
		IsProcessed: queryResult.IsProcessed,
	})
}

func insert(c *gin.Context) {
	requestBody := PasswordDto{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//validate if string consists of all spaces, etc
	preconditions := validateInput(requestBody.Password)
	isValid := arePreconditionsSatisfied(preconditions)
	//why do I use 400 in that case: https://stackoverflow.com/questions/3290182/which-status-code-should-i-use-for-failed-validations-or-invalid-duplicates
	if !isValid {
		c.JSON(http.StatusBadRequest, PostResponseDto{
			Id:            nil,
			Preconditions: preconditions,
			IsProcessed:   false,
		})
		return
	}
	filter := bson.D{{"password", requestBody.Password}}
	update := bson.D{{"$set", bson.D{{"password", requestBody.Password}, {"strength", -1}, {"isProcessed", false}}}}
	opts := options.Update().SetUpsert(true)
	result, err := PasswordCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		//mongo query issue - unlikely to happen
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//TODO id+preconditions
	c.JSON(http.StatusCreated, PostResponseDto{
		Id:            result.UpsertedID,
		Preconditions: preconditions,
		IsProcessed:   false,
	})
}

const lengthCondition = "Password should contain at least 4 signs"
const nonEmptyStringCondition = "Provided string is not empty"
const nonWhitespacesOnlyCondition = "Provided string doesnt consists of whitespaces only"
const notAllowedSymbolsCondition = "Provided string doesnt contain symbols that are not allowed"

func validateInput(password string) map[string]bool {

	var preconditions = map[string]bool{
		lengthCondition:             false,
		nonEmptyStringCondition:     false,
		nonWhitespacesOnlyCondition: false,
		notAllowedSymbolsCondition:  false}

	if len(password) >= 4 {
		preconditions[lengthCondition] = true
	}
	if !stringUtils.IsBlank(password) {
		preconditions[nonEmptyStringCondition] = true
	}
	if !stringUtils.IsWhitespace(password) {
		preconditions[nonWhitespacesOnlyCondition] = true
	}
	if !stringUtils.ContainsAny(password, "\\0") {
		preconditions[notAllowedSymbolsCondition] = true
	}

	return preconditions
}

func arePreconditionsSatisfied(preconditions map[string]bool) bool {
	for _, v := range preconditions {
		if v == false {
			return false
		}
	}
	return true
}

func updateStrength(c *gin.Context) {
}
