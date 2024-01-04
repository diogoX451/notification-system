package providers

import (
	"notification-system/internal/config/gorm"
	"notification-system/internal/database/migrate"
	"os"

	gormDB "gorm.io/gorm"
)

func ProviderConfig() *gorm.ConfigGorm {
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

	return config
}

func ProviderDB(c *gorm.ConfigGorm) (*gormDB.DB, error) {

	db, err := c.Connect()
	if err != nil {
		panic(err)
	}

	return db, nil
}
