package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/services"
)

type AccountHandler struct {
	accountService services.AccountService
}

func NewaccountHandler(accountService services.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

func (h *AccountHandler) DiscoverAccount(c *gin.Context) {
	var DiscoverAccountReq discoverAccountReq
	if err := c.ShouldBindJSON(&DiscoverAccountReq); err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	err := h.accountService.Registeraccount(c.Request.Context(), DiscoverAccountReq.Phone, DiscoverAccountReq.BankCode)
	c.JSON(200, gin.H{"error": err})
}

type discoverAccountReq struct {
	Phone    string `json:"phone" binding:"required"`
	BankCode string `json:"bank_code" binding:"required"`
}
