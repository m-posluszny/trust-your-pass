package configs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var passwordCollection *mongo.Collection
var mongoClient *mongo.Client

func InitMongoConnection() (*mongo.Client, *mongo.Collection) {
	InitEnvConfigs()
	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(EnvList.MongoUri))
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	PasswordCollection := Client.Database(EnvList.DbName).Collection(EnvList.CollectionName)

	log.Printf("Succesfully connected to mongo database")
	passwordCollection = PasswordCollection
	mongoClient = Client
	return Client, PasswordCollection
}

func GetMongoClient() *mongo.Client {
	if mongoClient == nil {
		InitMongoConnection()
	}
	return mongoClient
}

func GetPasswordCollection() *mongo.Collection {
	if passwordCollection == nil {
		InitMongoConnection()
	}
	return passwordCollection
}
