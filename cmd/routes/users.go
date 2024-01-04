package routes

import (
	"notification-system/internal/handlers"
	"notification-system/internal/repositories"
	"notification-system/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRouter(db *gorm.DB, rt *gin.RouterGroup) {
	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	rt.POST("/users-create", handler.Create)
}
