package main

import (
	"notification-system/cmd/routes"
	"notification-system/internal/config/gorm"
	"notification-system/internal/database/migrate"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Not exist .env")
	}
}

func main() {

	config := gorm.NewConfigGorm(
		gorm.WithHost(os.Getenv("DB_HOST")),
		gorm.WithPort(os.Getenv("DB_PORT")),
		gorm.WithUser(os.Getenv("DB_USER")),
		gorm.WithDBName(os.Getenv("DB_NAME")),
		gorm.WithPassword(os.Getenv("DB_PASSWORD")),
		gorm.WithSSLMode(os.Getenv("DB_SSL_MODE")),
		gorm.WithDriver(os.Getenv("DB_DRIVER")),
		gorm.WithDebug(true),
		gorm.WithMigrate(
			[]interface{}{
				&migrate.Groups{},
				&migrate.Companys{},
				&migrate.Users{},
				&migrate.NotificationsMethod{},
				&migrate.NotificationsDelivery{},
				&migrate.NotificationsMessage{},
				&migrate.NotificationsSent{},
			},
		),
	)

	db, err := config.Connect()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(gin.Logger())

	v1 := router.Group("/api/v1")

	routes.UsersRouter(db, v1)
	routes.GroupRoutes(db, v1)
	routes.CompanyRoutes(db, v1)

	router.Run(":8080")

}
