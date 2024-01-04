package migrate

import "gorm.io/gorm"

type NotificationsSent struct {
	gorm.Model
	ID                     int64                `gorm:"primaryKey;autoIncrement"`
	NotificationsMessageID int64                `gorm:"type:int;not null"`
	DateSent               int64                `gorm:"type:int;not null"`
	Status                 int                  `gorm:"type:int;not null"`
	NotificationsMessage   NotificationsMessage `gorm:"foreignKey:NotificationsMessageID"`
}
