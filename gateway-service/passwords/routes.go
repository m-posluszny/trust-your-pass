package passwords

import (
	"context"
	"inf/gateway-service/configs"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/v1/passwords", getAll)
	router.GET("/api/v1/passwords/:id", getById)
	router.POST("/api/v1/passwords", insert)

	//router.PATCH("/api/v1/passwords/:id", updateStrength)
}

func getAll(c *gin.Context) {
	//get all records that were processed
	cursor, err := configs.GetPasswordCollection().Find(c, bson.M{"strength": bson.M{"$ne": -1}})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}
	var results []PasswordDto
	if err = cursor.All(context.TODO(), &results); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(http.StatusOK, results)
}

func getById(c *gin.Context) {

	id := c.Param("id")

	var queryResult PasswordDto
	objId, _ := primitive.ObjectIDFromHex(id)
	err := configs.GetPasswordCollection().FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&queryResult)
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
	collection := configs.GetPasswordCollection()

	requestBody := PasswordDto{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	preconditions := validateInput(requestBody.Password)
	isValid := arePreconditionsSatisfied(preconditions)

	if !isValid {
		c.JSON(http.StatusOK, PostResponseDto{
			Id:            nil,
			Preconditions: preconditions,
			Strength:      -1,
			IsProcessed:   false,
		})
		return
	}

	filter := bson.D{{"password", requestBody.Password}}
	update := bson.D{{"$set", bson.D{{"password", requestBody.Password}, {"strength", -1}, {"isProcessed", false}}}}
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		//mongo query issue - unlikely to happen
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var queryResult PasswordDto
	//gets id of object if it was not upserted - avoids sending id with nil value
	if result.UpsertedID == nil {
		collection.FindOne(context.TODO(), filter).Decode(&queryResult)
		result.UpsertedID = queryResult.Id
	}

	PublishMessage(ModelInMessageDto{
		Id:       result.UpsertedID,
		Password: requestBody.Password,
	})

	//TODO id+preconditions
	c.JSON(http.StatusCreated, PostResponseDto{
		Id:            result.UpsertedID,
		Preconditions: preconditions,
		Strength:      -1,
		IsProcessed:   false,
	})
}

/*func updateStrength(c *gin.Context) {
	collection := configs.GetPasswordCollection()
	objId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var requestBody map[string]int
	if err := c.BindJSON(&requestBody); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", bson.D{{"strength", requestBody["strength"]}, {"IsProcessed", true}}}}
	opts := options.Update()
	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		//mongo query issue - unlikely to happen
		//no content?
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, bson.M{"updated": true})
}*/

func updateStrength(dto ModelOutMessageDto) {
	collection := configs.GetPasswordCollection()
	objId, _ := primitive.ObjectIDFromHex(dto.Id)
	filter := bson.D{{"_id", objId}}
	update := bson.D{{"$set", bson.D{{"strength", strconv.Itoa(dto.Strength)},
		{"isProcessed", false}}}}
	opts := options.Update()
	_, err := collection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		//mongo query issue - unlikely to happen
		log.Printf(err.Error())
		return
	}
}
