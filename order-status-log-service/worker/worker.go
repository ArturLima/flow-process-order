package worker

import (
	"log"
	"order-status-log-service/infra/rabbitmq"
)

type IWorker interface {
	StartWorker()
}

type Worker struct {
	consumer rabbitmq.IConsumer
}

func NewWorker() IWorker {
	return &Worker{
		consumer: rabbitmq.NewConsumer(),
	}
}

func (w *Worker) StartWorker() {
	messages, err := w.consumer.Delivery()

	if err != nil {
		log.Fatal("")
	}

	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages...")

	for {
		go func() {
			for message := range messages {
				log.Printf("[INFO] received event: %s\n", message.Body)
				message.Ack(false)
			}
		}()

		if <-w.consumer.NotifyClose() != nil {
			log.Printf("reconnecting with rabbitmq...")
			w.consumer.ReconnectConsumer()
			messages, err = w.consumer.Delivery()
		}
	}

}
