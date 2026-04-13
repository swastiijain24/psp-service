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
		log.Print("14")
		msg, err := w.consumer.Reader.FetchMessage(ctx)
		if err != nil {
			log.Printf("error fetching message: %v", err)
			continue
		}

		var paymentResponse pb.PaymentResponse

		err = proto.Unmarshal(msg.Value, &paymentResponse)
		if err != nil {
			log.Printf("error unpacking message: %v", err)
			continue
		}

		w.paymentService.PayentResponse(ctx, paymentResponse.GetTransactionId(), paymentResponse.GetStatus(), paymentResponse.GetDebitBankRef(), paymentResponse.GetCreditBankRef(), paymentResponse.GetFailureReason(), paymentResponse.GetCompletedAt())

		if err := w.consumer.Reader.CommitMessages(ctx, msg); err != nil {
			log.Printf("failed to commit: %v", err)
		}
		log.Print("16")
	}

}
