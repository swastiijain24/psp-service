package services

import (
	"context"
	"github.com/swastiijain24/psp/internals/gen"
	"github.com/swastiijain24/psp/internals/kafka"
	"google.golang.org/protobuf/proto"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context,transactionID string, payerVpa string, payeeVpa string, amount int64)
}

type PaymentSvc struct {
	vpaService         *VpaService
	paymentReqProducer *kafka.Producer
}

func NewPaymentService(vpaService *VpaService, paymentReqProducer *kafka.Producer) PaymentService {
	return &PaymentSvc{
		vpaService:         vpaService,
		paymentReqProducer: paymentReqProducer,
	}

}

func (s *PaymentSvc) ProcessPayment(ctx context.Context, transactionId string, payerVpa string, payeeVpa string, amount int64) {

	payerAccountID := s.vpaService.ResolveVpa(ctx, payerVpa)
	payeeAccountID := s.vpaService.ResolveVpa(ctx, payeeVpa)

	message := &pb.PaymentRequest{
		TransactionId:  transactionId,
		PayerAccountId: payerAccountID,
		PayeeAccountId: payeeAccountID,
		Amount:         amount,
	}

	data, err := proto.Marshal(message)
	if err != nil {
		return //err
	}

	s.paymentReqProducer.ProduceEvent(ctx, transactionId, data)

}
