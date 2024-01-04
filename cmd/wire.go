//go:build wireinject
// +build wireinject

package main

import (
	"notification-system/cmd/routes"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeRouter(db *gorm.DB, engine *gin.Engine) (*gin.RouterGroup, error) {
	wire.Build(
		routes.NewRouterGroup,
		routes.NewCompanyRoutesParams,
		routes.CompanyRoutes,
		routes.NewCompanyRoutesParams,
		routes.GroupRoutes,
	)
	return nil, nil

}

func NewRouterGroup(engine *gin.Engine) *gin.RouterGroup {
	return engine.Group("/api/v1")
}
