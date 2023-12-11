package configs

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var rmqConnection *amqp.Connection
var rmqChannel *amqp.Channel
var rmqModelInQueue *amqp.Queue
var rmqModelOutQueue *amqp.Queue

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
	GetModelInQueue()
	GetModelOutQueue()
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

func DeclareQueue(queueName string) *amqp.Queue {
	queue, err := GetChannel().QueueDeclare(
		queueName,
		true,
		false,
		false,
		true,
		nil,
	)
	if err != nil {
		log.Fatalf("Could not declare %s queue", queueName)
	}
	return &queue
}

func GetModelInQueue() *amqp.Queue {
	if rmqModelInQueue == nil {
		rmqModelInQueue = DeclareQueue(EnvList.RabbitMQModelInQueueName)
	}
	return rmqModelInQueue
}

func GetModelOutQueue() *amqp.Queue {
	if rmqModelOutQueue == nil {
		rmqModelOutQueue = DeclareQueue(EnvList.RabbitMQModelOutQueueName)
	}
	return rmqModelOutQueue
}
