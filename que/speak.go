package que

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func Speak(msg amqp.Publishing) {

	err := ConnChannel.Publish(
		"",       // exchange
		Que.Name, // routing key
		false,    // mandatory
		false,    // immediate
		msg,
	)

	failOnError(err, "Failed to Speak")

}
