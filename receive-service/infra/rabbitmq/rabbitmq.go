package rabbitmq

import (
	"errors"
	"os"
	"receive-service/utils"
	"time"

	"github.com/rabbitmq/amqp091-go"
	amqp "github.com/rabbitmq/amqp091-go"
)

type IRabbitMQ interface {
	Connection() *amqp.Connection
	Channel() *amqp091.Channel
	NotifyClose() <-chan error
	Reconnect()
}

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp091.Channel
	queue   amqp.Queue
	err     chan error
}

func NewRabbitMQ() IRabbitMQ {
	rabbitmq := &RabbitMQ{
		err: make(chan error),
	}

	rabbitmq.setConnection()
	rabbitmq.setQueue()

	return rabbitmq
}

func (r *RabbitMQ) Connection() *amqp.Connection {
	return r.conn
}

func (r *RabbitMQ) Channel() *amqp091.Channel {
	return r.setChannel()
}

func (r *RabbitMQ) NotifyClose() <-chan error {
	return r.err
}

func (r *RabbitMQ) Reconnect() {
	r.setConnection()
	r.setQueue()
}

func (r *RabbitMQ) setConnection() {
	amqpConfig := amqp.Config{
		Vhost:     "/",
		Heartbeat: 10 * time.Second,
		Locale:    "en_US",
	}

	c, err := amqp.DialConfig(os.Getenv("AMQP_SERVER_URL"), amqpConfig)
	if err != nil {
		utils.FailWithError("failure connecting to RabbitMQ", err)
	}

	go func() {
		<-r.conn.NotifyClose(make(chan *amqp.Error))
		r.err <- errors.New("Connection Closed")
	}()

	r.conn = c
}

func (r *RabbitMQ) setChannel() (c *amqp091.Channel) {
	c, err := r.conn.Channel()
	if err != nil {
		utils.FailWithError("failure creating channe", err)
		return
	}

	c.Confirm(false)

	err = c.Qos(1, 0, false)
	if err != nil {
		utils.FailWithError("failure to set Qos", err)
		return
	}

	return
}

func (r *RabbitMQ) setQueue() {
	q, _ := r.Channel().QueueDeclare(
		os.Getenv("AMQP_QUEUE"),
		true,
		false,
		false,
		false,
		nil)

	r.queue = q
}
