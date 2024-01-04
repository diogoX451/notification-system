package handlers

import (
	"fmt"
	dtos "notification-system/internal/dtos/user"
	"notification-system/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(service services.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var input dtos.CreateUserDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input)

	user, err := h.service.Create(input)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}
