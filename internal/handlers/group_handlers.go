package handlers

import (
	dtos "notification-system/internal/dtos/group"
	"notification-system/internal/services"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	service services.IGroupService
}

func NewGroupHandler(service services.IGroupService) *GroupHandler {
	return &GroupHandler{service}
}

func (h *GroupHandler) Create(ctx *gin.Context) {
	var input dtos.CreateGroupDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	group, err := h.service.Create(&input)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, group)
}
