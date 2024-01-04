package providers

import (
	"github.com/gin-gonic/gin"
)

func ProviderGinEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	return router
}

func ProviderRouterGroup(router *gin.Engine, prefix string) *gin.RouterGroup {
	group := router.Group(prefix)
	return group
}
