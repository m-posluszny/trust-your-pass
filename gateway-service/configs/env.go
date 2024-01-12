package configs

import (
	"log"

	"github.com/spf13/viper"
)

var EnvList *envList

func InitEnvConfigs() {
	EnvList = loadEnvVariables()
}

// struct to map env values
type envList struct {
	ServerAddress             string `mapstructure:"SERVER_ADDRESS"`
	MongoUri                  string `mapstructure:"MONGO_URI"`
	DbName                    string `mapstructure:"MONGO_DB_NAME"`
	CollectionName            string `mapstructure:"MONGO_COLLECTION_NAME"`
	RabbitMQConnectionString  string `mapstructure:"RABBITMQ_CONNECTION_STRING"`
	RabbitMQModelInQueueName  string `mapstructure:"RABBITMQ_MODEL_IN_QUEUE_NAME"`
	RabbitMQModelOutQueueName string `mapstructure:"RABBITMQ_MODEL_OUT_QUEUE_NAME"`
}

func loadEnvVariables() (config *envList) {
	viper.AddConfigPath(".")
	viper.SetConfigName("application")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
