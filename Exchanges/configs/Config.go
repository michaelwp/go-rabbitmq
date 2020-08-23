package configs

import "os"

func SetRabbitUrl() {
	_ = os.Setenv(
		"rabbitMqTes",
		"amqp://gsvcrhuy:XI3kxXKAr1NoKM1A7GIW8E0uevVfKIKb@coyote.rmq.cloudamqp.com/gsvcrhuy",
		)
}
