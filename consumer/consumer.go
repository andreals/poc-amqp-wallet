package main

import (
	"fmt"

	"github.com/andreals/poc-amqp-wallet/consts"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer RabbitMQ")

	chann := make(chan bool)

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

	msgs, err := ch.Consume(consts.WALLET_QUEUE, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		var counter int64

		for msg := range msgs {
			fmt.Printf("Message (%d): %s\n", counter, string(msg.Body))
			counter++
		}
	}()

	fmt.Println("Connected to RabbitMQ instance")
	<-chann

}
