package que

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Host string = "amqp://guest:guest@localhost:5672/"

	Conn *amqp.Connection

	ConnChannel *amqp.Channel

	Que amqp.Queue
)

func ConnectService() {
	con, err := amqp.Dial(Host)
	failOnError(err, "Failed to connect to RabbitMQ")

	Conn = con
}

func DisConnectService() {

	Conn.Close()

}

func ConnectChannel() {

	ConnChannel, _ = Conn.Channel()

}

func DisConnectChanel() {

	ConnChannel.Close()

}

func CreateQue(QueName string) {
	Que, _ = ConnChannel.QueueDeclare(
		QueName,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
