package configs

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rmqConnection *amqp.Connection
var rmqChannel *amqp.Channel

func InitRabbitMQConnection() (*amqp.Connection, *amqp.Channel) {
	InitEnvConfigs()

	conn, err := amqp.Dial(EnvList.RabbitMQConnectionString)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	log.Printf("Succesfully connected to rabbitmq broker")
	rmqConnection = conn
	rmqChannel = ch
	return conn, ch
}

func GetChannel() *amqp.Channel {
	if rmqChannel == nil {
		InitRabbitMQConnection()
	}
	return rmqChannel
}

func GetConnection() *amqp.Connection {
	if rmqConnection == nil {
		InitRabbitMQConnection()
	}
	return rmqConnection
}
