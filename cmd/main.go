package main

import (
	"notification-system/cmd/routes"
	"notification-system/cmd/wire"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Not exist .env")
	}
}

func main() {
	app, err := wire.InitializeEvent("/api/v1")
	if err != nil {
		panic(err)
	}

	app.RouterGroup.Use(gin.Logger())

	routes.CompanyRoutes(app.DB, app.RouterGroup)
	routes.GroupRoutes(app.DB, app.RouterGroup)
	routes.UsersRouter(app.DB, app.RouterGroup)

}
