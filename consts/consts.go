package consts

import (
	"fmt"
)

const (
	LOGIN_RABBITMQ    = "guest"
	PASSWORD_RABBITMQ = "guest"
	HOST_RABBITMQ     = "localhost"
	PORT_RABBITMQ     = "5672"
	WEBHOOK_QUEUE     = "WebhookQueue"
	WALLET_QUEUE      = "WalletQueue"
)

var (
	URLRabbitMQ = fmt.Sprintf("amqp://%s:%s@%s:%s/", LOGIN_RABBITMQ, PASSWORD_RABBITMQ, HOST_RABBITMQ, PORT_RABBITMQ)
)
