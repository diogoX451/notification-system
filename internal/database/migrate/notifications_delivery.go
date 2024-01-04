package migrate

import "gorm.io/gorm"

type NotificationsDelivery struct {
	gorm.Model
	ID                   int64               `gorm:"primaryKey;autoIncrement"`
	Recived              bool                `gorm:"type:boolean;not null"`
	NotificationMethodID int64               `gorm:"type:varchar(255);not null"`
	UserID               int64               `gorm:"type:varchar(255);not null"`
	NotificationMethod   NotificationsMethod `gorm:"foreignKey:NotificationMethodID"`
	Users                Users               `gorm:"foreignKey:UserID"`
}
