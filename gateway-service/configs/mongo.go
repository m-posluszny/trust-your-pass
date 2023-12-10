package configs

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

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
	return Client, PasswordCollection
}
