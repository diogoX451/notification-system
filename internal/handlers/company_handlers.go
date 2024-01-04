package handlers

import (
	dtos "notification-system/internal/dtos/company"
	"notification-system/internal/services"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service services.ICompanyService
}

func NewCompanyHandler(service services.ICompanyService) *CompanyHandler {
	return &CompanyHandler{service: service}
}

func (h *CompanyHandler) Create(ctx *gin.Context) {
	var input dtos.CreateCompanyDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	message, err := h.service.Create(input)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, message)
}
