package configs

import (
	"log"

	"github.com/spf13/viper"
)

// Initilize this variable to access the env values
var EnvList *envList

// We will call this in main.go to load the env variables
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

// Call to load the variables from env
func loadEnvVariables() (config *envList) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("../tools/")

	// Tell viper the name of your file
	viper.SetConfigName("services")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
