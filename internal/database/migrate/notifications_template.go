package migrate

import "gorm.io/gorm"

type NotificationsTemplate struct {
	gorm.Model
	ID int64 `gorm:"primaryKey;autoIncrement"`
	
}
