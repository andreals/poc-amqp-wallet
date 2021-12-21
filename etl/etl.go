package main

import (
	"encoding/json"
	"fmt"

	"github.com/andreals/poc-amqp-wallet/consts"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("ETL RabbitMQ")

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

	msgs, err := ch.Consume(consts.WEBHOOK_QUEUE, "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	q, err := ch.QueueDeclare(consts.WALLET_QUEUE, false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(q)

	go func() {
		for msg := range msgs {
			fmt.Printf("Message: %s\n", string(msg.Body))
			var contents []map[string]string

			err := json.Unmarshal(msg.Body, &contents)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, content := range contents {

				contentMarshal, err := json.Marshal(content)
				if err != nil {
					fmt.Println(err)
					return
				}

				err = ch.Publish("", consts.WALLET_QUEUE, false, false, amqp.Publishing{
					ContentType: "application/json",
					Body:        contentMarshal,
				})
				if err != nil {
					fmt.Println(err)
					return
				}
			}

		}
	}()

	fmt.Println("Connected to RabbitMQ instance")
	<-chann

}
