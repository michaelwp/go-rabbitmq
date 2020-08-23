package main

/**
Rabbit MQ Consumers
Created by MPutong, 20082020
*/

import (
	"github.com/michaelwp/go-rabbitmq/configs"
	"github.com/michaelwp/go-rabbitmq/errHandlers"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// set rabbitMq url env
	configs.SetRabbitUrl()

	// connect to rabbitmq server
	conn, err := amqp.Dial(os.Getenv("rabbitMqTes"))
	// if error call failonerror function
	errHandlers.FailOnError(err, "Failed to connect to RabbitMQ")
	// close connection at the end
	defer conn.Close()

	// open channel
	ch, err := conn.Channel()
	// if error call failonerror function
	errHandlers.FailOnError(err, "Failed to open a channel")
	// close channel at the end
	defer ch.Close()

	/* declare queue
	------------------------------------------------------------------------------------------------------------------
	 (1) A true durable queue means that the queue definition will survive a server restart, not the messages in it
	 (2) A true exclusive queue can only be used (consumed from, purged, deleted, etc) by its declaring connection.
	     An attempt to use an exclusive queue from a different connection will result in a channel-level exception
	     RESOURCE_LOCKED with an error message that says cannot obtain exclusive access to locked queue.
	 (3) The client should not wait for a reply method. If the server could not complete the method it will raise
	     a channel or connection exception.
	*/
	q, err := ch.QueueDeclare(
		"hello", // queue name/ routing key
		false, // durable (1)
		false, // delete when unused
		false, // exclusive (2)
		false, // no-wait (3)
		nil, // arguments
	)
	// if error call failonerror function
	errHandlers.FailOnError(err, "Failed to declare a queue")

	// consume the messages
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	// if error call failonerror function
	errHandlers.FailOnError(err, "Failed to register a consumer")

	// create channel
	forever := make(chan bool)

	// create goroutine
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	// print waiting messages
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// call channel
	<-forever
}