package migrate

import "gorm.io/gorm"

type NotificationsMethod struct {
	gorm.Model
	ID          int64  `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:varchar(255);not null"`
}
