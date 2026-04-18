package workers

import (
	"context"
	"log"
	"github.com/swastiijain24/psp/internals/kafka"
	pb "github.com/swastiijain24/psp/internals/pb"
	"github.com/swastiijain24/psp/internals/services"
	"google.golang.org/protobuf/proto"
)

type ResponseWorker struct {
	consumer       *kafka.Consumer
	paymentService services.PaymentService
}

func NewResponseWorker(consumer *kafka.Consumer, paymentService services.PaymentService) *ResponseWorker {
	return &ResponseWorker{
		consumer:       consumer,
		paymentService: paymentService,
	}
}

func (w *ResponseWorker) StartConsumingResponse(ctx context.Context) {

	for {
		msg, err := w.consumer.Reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil{
				log.Print("stopping consumer loop")
				return 
			} 
			log.Printf("error fetching message: %v", err)
			continue
		}

		log.Print("processing final response message from Kafka")

		var paymentResponse pb.PaymentResponse

		err = proto.Unmarshal(msg.Value, &paymentResponse)
		if err != nil {
			log.Printf("error unpacking message: %v", err)
			continue
		}
		if err = w.paymentService.PayentResponse(ctx, &paymentResponse); err != nil {
			log.Println(err.Error())
			continue 
		}

		if err := w.consumer.Reader.CommitMessages(ctx, msg); err != nil {
			log.Printf("failed to commit: %v", err)
			continue 
		}
		log.Print("processed final response")
	}

}
