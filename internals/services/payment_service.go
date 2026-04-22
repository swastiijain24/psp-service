package services

import (
	"context"
	"fmt"
	"log"

	"github.com/swastiijain24/psp/internals/kafka"
	"github.com/swastiijain24/psp/internals/pb"
	"github.com/swastiijain24/psp/internals/redis"
	"google.golang.org/protobuf/proto"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, transactionID string, payerVpa string, payeeVpa string, amount int64, mpin string) error
	PayentResponse(ctx context.Context, paymentResponse *pb.PaymentResponse) error
	GetTransactionStatus(ctx context.Context, transactionId string) (string, error)
}

type PaymentSvc struct {
	vpaService         VpaService
	paymentReqProducer *kafka.Producer
	redis              *repository.RedisStore
}

func NewPaymentService(vpaService VpaService, paymentReqProducer *kafka.Producer, redis *repository.RedisStore) PaymentService {
	return &PaymentSvc{
		vpaService:         vpaService,
		paymentReqProducer: paymentReqProducer,
		redis:              redis,
	}

}

func (s *PaymentSvc) ProcessPayment(ctx context.Context, transactionId string, payerVpa string, payeeVpa string, amount int64, mpin string) error {

	payerAccountID, payerBankCode, err := s.vpaService.ResolveVpa(ctx, payerVpa)
	if err != nil {
		return fmt.Errorf("invalid VPA ID :%w", err)
	}
	payeeAccountID, payeeBankCode, err := s.vpaService.ResolveVpa(ctx, payeeVpa)
	if err != nil {
		return fmt.Errorf("invalid VPA ID :%w", err)
	}

	message := &pb.PaymentRequest{
		TransactionId:  transactionId,
		PayerAccountId: payerAccountID,
		PayeeAccountId: payeeAccountID,
		Amount:         amount,
		PayerBankCode:  payerBankCode,
		PayeeBankCode:  payeeBankCode,
		Mpin: mpin,
	}

	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to pack data: %w", err)
	}

	err = s.redis.SetInitialStatus(ctx, transactionId)
	if err != nil {
		return fmt.Errorf("Failed to set initial status in Redis for %s: %v", transactionId, err)
	}

	log.Print("request set in redis 2")

	err = s.paymentReqProducer.ProduceEvent(ctx, transactionId, data, "payment.request.v1")
	if err != nil {
		log.Printf("Error producing payment event for %s: %v", transactionId, err)
		_ = s.redis.DeleteStatus(ctx, transactionId)
		return fmt.Errorf("Deleted entry in Redis for %s: %v", transactionId, err)

	}
	log.Print("payment request produced from psp service 3")
	return nil

}

func (s *PaymentSvc) PayentResponse(ctx context.Context, paymentResponse *pb.PaymentResponse) error {
	
	err := s.redis.UpdateFinalStatus(ctx, paymentResponse.GetTransactionId(), paymentResponse.GetStatus())
	if err != nil {
		return fmt.Errorf("Failed to update final status in Redis for %s: %v", paymentResponse.GetTransactionId(), err)
	}
	log.Print("transaction donee")
	return nil

}

func (s *PaymentSvc) GetTransactionStatus(ctx context.Context, transactionId string) (string, error) {
	return s.redis.GetStatus(ctx, transactionId)
}
