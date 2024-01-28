package rabbitmq

import (
	"context"
	"encoding/json"
	"os"
	"receive-service/core/models"

	"github.com/rabbitmq/amqp091-go"
)

type IPublisher interface {
	SendToOrderStatus(o *models.OrderStatus) (err error)
}

type Publisher struct {
	rabbitmq IRabbitMQ
}

func NewPublisher() IPublisher {
	return &Publisher{
		rabbitmq: NewRabbitMQ(),
	}
}

func (p *Publisher) SendToOrderStatus(o *models.OrderStatus) (err error) {
	m, _ := json.Marshal(o)
	message := amqp091.Publishing{
		ContentType: "application/json",
		Body:        []byte(m),
	}

	err = p.rabbitmq.Channel().PublishWithContext(
		context.Background(),
		"",                                       // exchange
		os.Getenv("AMQP_QUEUE_STATUS_PUBLISHER"), // queue name
		false,                                    // mandatory
		false,                                    // immediate
		message,                                  // message to publish
	)

	return
}
