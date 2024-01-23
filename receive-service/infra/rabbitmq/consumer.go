package rabbitmq

import (
	"os"

	"github.com/rabbitmq/amqp091-go"
)

type IConsumer interface {
	Delivery() (msgs <-chan amqp091.Delivery, err error)
	NotifyClose() <-chan error
	ReconnectConsumer()
}

type Consumer struct {
	rabbitmq IRabbitMQ
}

func NewConsumer() IConsumer {
	return &Consumer{
		rabbitmq: NewRabbitMQ(),
	}
}

func (c *Consumer) Delivery() (msgs <-chan amqp091.Delivery, err error) {
	if c.rabbitmq.Channel() == nil {
		return
	}

	msgs, err = c.rabbitmq.Channel().Consume(
		os.Getenv("AMQP_QUEUE"), // queue name
		"receive-service",       // consumer
		false,                   // auto-ack
		false,                   // exclusive
		false,                   // no local
		false,                   // no wait
		nil)

	return
}

func (c *Consumer) NotifyClose() <-chan error {
	return c.rabbitmq.NotifyClose()
}

func (c *Consumer) ReconnectConsumer() {
	c.rabbitmq.Reconnect()
}
