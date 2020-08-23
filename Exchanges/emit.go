package main

import (
	"github.com/michaelwp/go-rabbitmq/configs"
	"github.com/michaelwp/go-rabbitmq/errHandlers"
	"github.com/michaelwp/go-rabbitmq/helpers"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func main() {
	// set rabbitMq url env
	configs.SetRabbitUrl()

	conn, err := amqp.Dial(os.Getenv("rabbitMqTes"))
	errHandlers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	errHandlers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	errHandlers.FailOnError(err, "Failed to declare an exchange")

	body := helpers.BodyFrom(os.Args)
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(body),
		})
	errHandlers.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}