package routes

import (
	"notification-system/internal/handlers"
	"notification-system/internal/repositories"
	"notification-system/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GroupRoutes(db *gorm.DB, rt *gin.RouterGroup) {
	repo := repositories.NewGroupRepository(db)
	service := services.NewGroupService(repo)
	handler := handlers.NewGroupHandler(service)

	rt.POST("/groups-create", handler.Create)
}
