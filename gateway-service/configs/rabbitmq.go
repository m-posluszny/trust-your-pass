package configs

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

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
	return conn, ch
}
