package main

/**
Rabbit MQ Publishers
Created by MPutong, 20082020
 */

import (
	"github.com/michaelwp/go-rabbitmq/configs"
	"github.com/michaelwp/go-rabbitmq/errHandlers"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
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
	errHandlers.FailOnError(err, "Failed to open channel")
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

	// set the body messages will send
	body := "Hello World"

	/* publish the body messages
	------------------------------------------------------------------------------------------------------------------
	1. Exchanges are message routing agents, defined by the virtual host within RabbitMQ. An exchange is responsible
	   for routing the messages to different queues with the help of header attributes, bindings, and routing keys.
	   A binding is a "link" that you set up to bind a queue to an exchange.
	2. When a published message cannot be routed to any queue (e.g. because there are no bindings defined for the
	   target exchange), and the publisher set the mandatory message property to false (this is the default),
	   the message is discarded or republished to an alternate exchange, if any.
	   When a published message cannot be routed to any queue, and the publisher set the mandatory message property
	   to true, the message will be returned to it.
	3. For a message published with immediate set. If there is at least one consumer connected to my queue that
	   can take delivery of a message right this moment, deliver this message to them immediately. If there are no
	   consumers connected then there's no point in having my message consumed later and they'll never see it.
	 */
	err = ch.Publish(
		"",     // exchange (1)
		q.Name, // queue name/ routing key
		false,  // mandatory (2)
		false,  // immdiate (3)
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
		},
	)
	// if error call failonerror function
	errHandlers.FailOnError(err, "Failed to publish a message")
	// if succeed to sent the body message
	log.Printf(" [x] Sent %s", body)
}