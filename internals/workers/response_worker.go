package workers

import (
	"context"

	"github.com/swastiijain24/psp/internals/kafka"
)

type ResponseWorker struct {
	consumer *kafka.Consumer
}

func NewResponseWorker(consumer *kafka.Consumer) *ResponseWorker {
	return &ResponseWorker{
		consumer: consumer,
	}
}

func (w *ResponseWorker) StartConsumingResponse(ctx context.Context) {

	for {
		
		msg , err := w.consumer.Reader.ReadMessage(ctx)
		if err != nil {
			break
		}

		
	}

}