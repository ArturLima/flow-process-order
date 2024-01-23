package worker

import (
	"log"
	"receive-service/core/handlers"
	"receive-service/infra/rabbitmq"
)

type IWorker interface {
	StartWorker()
}

type Worker struct {
	consumer rabbitmq.IConsumer
	handler  handlers.IOrderHandler
}

func NewWorker() IWorker {
	return &Worker{
		consumer: rabbitmq.NewConsumer(),
		handler:  handlers.NewOrderHandler(),
	}
}

func (w *Worker) StartWorker() {
	messages, err := w.consumer.Delivery()
	if err != nil {
		log.Fatal("")
	}
	for {
		go func() {
			for message := range messages {
				log.Printf("[INFO] received event: %s\n", message.Body)
				if done := w.handler.Execute(message.Body); done {
					message.Ack(false)
				}
			}
		}()

		if <-w.consumer.NotifyClose() != nil {
			log.Printf("reconnecting with rabbitmq...")
			w.consumer.ReconnectConsumer()
			messages, err = w.consumer.Delivery()
		}
	}

}
