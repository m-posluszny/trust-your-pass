package passwords

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"inf/gateway-service/configs"
	"log"
)

func PublishMessage(message interface{}) {

	encoded, err := json.Marshal(message)
	if err != nil {
		log.Fatal("Cannot marshall message")
	}

	err = configs.GetChannel().PublishWithContext(context.Background(), "",
		configs.EnvList.RabbitMQModelInQueueName, false, false, amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         encoded,
		})
	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
}

func ConsumeMessages(messageType interface{}) {

	messages, err := configs.GetChannel().ConsumeWithContext(context.Background(),
		configs.EnvList.RabbitMQModelOutQueueName, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error consuming message: %s", err)
	}

	outMessage := messageType

	k := make(chan bool)
	go func() {
		for d := range messages {
			err = json.Unmarshal(d.Body, &outMessage)
			if err != nil {
				log.Fatalf("Error decoding message: %s", err)
			}
		}
	}()
	<-k
}
