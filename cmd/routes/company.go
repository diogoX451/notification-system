package routes

import (
	"notification-system/internal/handlers"
	"notification-system/internal/repositories"
	"notification-system/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CompanyRoutes(db *gorm.DB, rt *gin.RouterGroup) {
	repo := repositories.NewCompanyRepository(db)
	service := services.NewCompanyService(repo)
	handler := handlers.NewCompanyHandler(service)

	rt.POST("/company-create", handler.Create)
}
