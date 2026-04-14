package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/services"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
	}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var params paymentParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	log.Print("received http payment request")

	err = h.paymentService.ProcessPayment(c.Request.Context(), params.TransactionID, params.PayerVPA, params.PayeeVPA, params.Amount, params.Mpin)
	if err != nil {
		log.Println(err)
	}

	log.Print("handler function's job done")
}

func (h *PaymentHandler) GetTxnStatus(c *gin.Context) {
	//psp will poll on this to get  the status of the txn

	transactionId := c.Param("transactionId")
	h.paymentService.GetTransactionStatus(c.Request.Context(), transactionId)
}

type paymentParams struct {
	TransactionID string `json:"transaction_id" binding:"required"`
	PayerVPA      string `json:"payer_vpa" binding:"required"`
	PayeeVPA      string `json:"payee_vpa" binding:"required"`
	Amount        int64  `json:"amount" binding:"required"`
	Mpin          string `json:"mpin" binding:"required"`
}
