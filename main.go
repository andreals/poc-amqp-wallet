package main

import (
	"fmt"

	"github.com/andreals/poc-amqp-wallet/consts"
	"github.com/streadway/amqp"
)

func main() {

	contentJSON := []byte("{'id': 12}")

	conn, err := amqp.Dial(consts.URLRabbitMQ)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(consts.WEBHOOK_QUEUE, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(q)

	err = ch.Publish("", consts.WEBHOOK_QUEUE, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        contentJSON,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully published message to queue")
}
