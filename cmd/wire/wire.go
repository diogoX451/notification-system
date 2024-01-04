//go:build wireinject
// +build wireinject

package wire

import (
	"notification-system/internal/providers"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type App struct {
	DB          *gorm.DB
	RouterGroup *gin.RouterGroup
}

func InitializeEvent(prefix string) (*App, error) {
	panic(
		wire.Build(
			providers.ProviderConfig,
			providers.ProviderDB,
			providers.ProviderGinEngine,
			providers.ProviderRouterGroup,
			wire.Struct(new(App), "*"),
		),
	)
}
