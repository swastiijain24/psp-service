package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swastiijain24/psp/internals/services"
)

type VpaHandler struct {
	vpaService services.VpaService
}

func NewVpaHandler(vpaService services.VpaService) *VpaHandler {
	return &VpaHandler{
		vpaService: vpaService,
	}
}

func (h *VpaHandler) RegisterVpa(c *gin.Context){
	var vpaRegisterReq VpaRegisterReq
	if err := c.ShouldBindJSON(&vpaRegisterReq); err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	err := h.vpaService.RegisterVpa(c.Request.Context(), vpaRegisterReq.VpaId, vpaRegisterReq.VpaId, vpaRegisterReq.BankCode)
	c.JSON(200, gin.H{"error":err})
}

type VpaRegisterReq struct {
	VpaId     string `json:"vpa_id" binding:"required"`
	AccountId string `json:"account_id" binding:"required"`
	BankCode  string `json:"bank_code" binding:"required"`
}