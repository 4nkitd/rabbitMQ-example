package main

import (
	"fmt"
	"rnoti/que"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	fmt.Println("app runing...")

	que.ConnectService()
	defer que.DisConnectService()

	que.CreateQue("task_queue")

	que.ConnectChannel()
	defer que.DisConnectService()

	msg := "hey!"

	msgBody := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msg),
	}

	que.Speak(msgBody)

	go que.Listen()

	time.Sleep(20 * time.Second)

	que.Speak(msgBody)

}
